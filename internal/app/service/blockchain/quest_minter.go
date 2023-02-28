package blockchain

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"go.uber.org/zap"
	"math/big"
	"strings"
	"time"
)

var questMinterAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.QuestMinterMetaData.ABI))
	if err != nil {
		panic(err)
	}
	questMinterAbi = contractAbi
}

func (b *BlockChain) handleClaimed(hash string, vLog *types.Log) (err error) {
	var claimed ABI.QuestMinterClaimed
	if err = questMinterAbi.UnpackIntoInterface(&claimed, "Claimed", vLog.Data); err != nil {
		return
	}
	// no such tokenId in quest
	id, err := b.dao.HasTokenId(vLog.Topics[0].Big().Uint64())
	if err != nil {
		return errors.New("no such tokenId in quest")
	}
	//
	challenges := model.UserChallenges{
		Address: claimed.Sender.String(),
		QuestID: id,
		Status:  2,
		Claimed: true,
		ClaimTs: time.Now().Unix(),
	}
	err = b.dao.CreateChallenges(&challenges)
	if err != nil {
		return err
	}
	b.handleTraverseStatus(hash, 1, "")
	return
}

func (b *BlockChain) AirdropBadge() error {
	log.Warn("AirdropBadge Run")
	client, err := ethclient.Dial(b.c.BlockChain.Provider)
	if err != nil {
		return errors.New("ethclient dial error")
	}
	paddingList, err := b.dao.GetPendingAirdrop()
	if err != nil {
		return err
	}
	for tokenId, list := range paddingList {
		receivers := b.receiverNotClaimList(client, tokenId, list)
		hash, err := b._airdropBadge(client, tokenId, receivers)
		if err != nil {
			log.Error("_airdropBadge", zap.Any("error", err))
			continue
		}
		if err := b.dao.UpdateAirdroppedList(tokenId, receivers, hash.String()); err != nil {
			log.Error("updateAirdropStatus", zap.Any("error", err))
		}
	}
	return nil
}

func (b *BlockChain) _airdropBadge(client *ethclient.Client, tokenID int64, receivers []common.Address) (txHash common.Hash, err error) {
	tokenId := big.NewInt(tokenID)

	privateKey, err := crypto.HexToECDSA(b.c.BlockChain.PrivateKey)
	if err != nil {
		return
	}
	address, err := utils.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"string", "uint256", "address", "address"},
		// values
		[]interface{}{
			"airdropBadge", tokenId, b.c.Contract.Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27

	questMinter, err := ABI.NewQuestMinter(common.HexToAddress(b.c.Contract.QuestMinter), client)
	if err != nil {
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(b.c.BlockChain.ChainID))
	if err != nil {
		return
	}
	transactOpts := &bind.TransactOpts{
		From:     address,
		Nonce:    auth.Nonce,
		Signer:   auth.Signer,
		Value:    big.NewInt(0),
		GasPrice: auth.GasPrice,
		GasLimit: auth.GasLimit,
		Context:  auth.Context,
		NoSend:   false,
	}
	tx, err := questMinter.AirdropBadge(transactOpts, tokenId, receivers, signature)
	if err != nil {
		return
	}
	log.Info("Airdrop tx sent :", zap.String("hash: ", tx.Hash().Hex()))
	return tx.Hash(), nil
}

func (b *BlockChain) receiverNotClaimList(client *ethclient.Client, tokenId int64, receivers []string) (receiversNotClaim []common.Address) {
	for _, receiver := range receivers {
		badge, err := ABI.NewBadge(common.HexToAddress(b.c.Contract.Badge), client)
		if err != nil {
			continue
		}
		res, err := badge.BalanceOf(nil, common.HexToAddress(receiver), big.NewInt(tokenId))
		if err != nil {
			continue
		}
		if res.Cmp(big.NewInt(0)) != 0 {
			// already claimed update status
			if err = b.dao.UpdateAirdropped(&model.ClaimBadgeTweet{Address: receiver, TokenId: tokenId}); err != nil {
				log.Error("UpdateAirdropped error", zap.Error(err))
			}
			continue
		}
		receiversNotClaim = append(receiversNotClaim, common.HexToAddress(receiver))
	}
	return
}
