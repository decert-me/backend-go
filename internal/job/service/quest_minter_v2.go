package service

import (
	ABIV2 "backend-go/abi/v2"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

var questMinterAbiV2 abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABIV2.BadgeMinterV2MetaData.ABI))
	if err != nil {
		panic(err)
	}
	questMinterAbiV2 = contractAbi
}

func (s *Service) AirdropBadgeV2() error {
	provider := s.w.Next()
	defer func() {
		if err := recover(); err != nil {
			provider.OnInvokeFault()
			log.Errorv("AirdropBadge error", zap.Any("error", err), zap.Any("provider", provider))
		}
	}()
	client, err := ethclient.Dial(provider.Item)
	if err != nil {
		log.Error("ethclient dial error")
		return errors.New("ethclient dial error")
	}
	tokenIds, listAddr, scores, err := s.dao.GetPendingAirdrop()
	if err != nil {
		log.Error("GetPendingAirdrop error")
		return err
	}
	if len(tokenIds) == 0 { // no task return
		return nil
	}
	tokenIdRes, receivers, scores := s.receiverNotClaimListV2(client, tokenIds, listAddr, scores)
	log.Warn("AirdropBadge Run")
	hash, err := s._airdropBadge(client, tokenIdRes, receivers, scores)
	if err != nil {
		log.Errorv("_airdropBadge", zap.Any("error", err))
		return nil
	}
	if err := s.dao.UpdateAirdroppedList(tokenIdRes, receivers, hash.String()); err != nil {
		log.Errorv("updateAirdropStatus", zap.Any("error", err))
	}
	if err := s.dao.CreateChallengesList(tokenIdRes, receivers, scores); err != nil {
		log.Errorv("updateAirdropStatus", zap.Any("error", err))
	}
	provider.OnInvokeSuccess()
	return nil
}

func (s *Service) _airdropBadgeV2(client *ethclient.Client, tokenIDs []*big.Int, receivers []common.Address, scores []*big.Int) (txHash common.Hash, err error) {
	signPrivateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	airdropPrivateKey, err := crypto.HexToECDSA(s.c.BlockChain.AirdropPrivateKey)
	if err != nil {
		return
	}
	airdropAddress, err := utils.PrivateKeyToAddress(airdropPrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"string", "uint256[]", "address[]", "address", "address"},
		// values
		[]interface{}{
			"airdropBadge", tokenIDs, receivers, s.c.Contract.V2.BadgeMinter, airdropAddress,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, signPrivateKey)
	signature[64] += 27

	badgeMinter, err := ABIV2.NewBadgeMinterV2(common.HexToAddress(s.c.Contract.V2.BadgeMinter), client)
	if err != nil {
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(airdropPrivateKey, big.NewInt(s.c.BlockChain.ChainID))
	if err != nil {
		return
	}
	transactOpts := &bind.TransactOpts{
		From:     airdropAddress,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    big.NewInt(0),
		GasPrice: auth.GasPrice,
		GasLimit: auth.GasLimit,
		Context:  auth.Context,
		NoSend:   false,
	}
	// TODO: uris
	var uris []string
	// 获取Quest MetaData
	for _, tokenID := range tokenIDs {
		quest, err := s.dao.GetQuestByTokenID(tokenID.Int64())
		if err != nil {
			log.Errorv("GetQuestByTokenID error", zap.Any("tokenID", tokenID), zap.Error(err))
			continue
		}
		// TODO
		// 处理Metadata

		// 上传到IPFS

		uris = append(uris, quest.Uri)
	}
	tx, err := badgeMinter.AirdropBadge(transactOpts, tokenIDs, receivers, uris, signature)
	if err != nil {
		log.Errorv("questMinter.AirdropBadge error", zap.Any("tokenIDs", tokenIDs), zap.Any("receivers", receivers), zap.Any("signature", signature), zap.Error(err))
		return
	}
	log.Infov("Airdrop tx sent :", zap.String("hash: ", tx.Hash().Hex()))
	return tx.Hash(), nil
}

func (s *Service) receiverNotClaimListV2(client *ethclient.Client, tokenId []*big.Int, receivers []string, scores []*big.Int) (tokenIdRes []*big.Int, receiversNotClaim []common.Address, scoresRes []*big.Int) {
	badge, err := ABIV2.NewBadgeV2(common.HexToAddress(s.c.Contract.V2.Badge), client)
	if err != nil {
		return
	}
	for i, _ := range receivers {
		if !utils.IsValidAddress(receivers[i]) {
			continue
		}
		res, err := badge.AddrToQuestToBadge(nil, common.HexToAddress(receivers[i]), tokenId[i])
		if err != nil {
			continue
		}
		if res.Cmp(big.NewInt(0)) != 0 {
			// already claimed update status
			if err = s.dao.UpdateAirdroppedError(tokenId[i].Int64(), receivers[i], "already claimed"); err != nil {
				log.Errorv("UpdateAirdropped error", zap.Error(err))
			}
			continue
		}
		tokenIdRes = append(tokenIdRes, tokenId[i])
		receiversNotClaim = append(receiversNotClaim, common.HexToAddress(receivers[i]))
		scoresRes = append(scoresRes, scores[i])
	}
	if len(tokenIdRes) != len(receiversNotClaim) {
		err = errors.New("token and address len error")
		return
	}
	return
}
