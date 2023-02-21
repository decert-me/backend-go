package blockchain

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"math/big"
	"strings"
	"time"
)

var questMinterAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.QuestMinterMetaData.ABI))
	if err != nil {
		global.LOG.Error("Failed to Load Abi", zap.Error(err))
		panic(err)
	}
	questMinterAbi = contractAbi
}

func handleClaimed(hash string, vLog *types.Log) (err error) {
	var claimed ABI.QuestMinterClaimed
	if err = questMinterAbi.UnpackIntoInterface(&claimed, "Claimed", vLog.Data); err != nil {
		return
	}
	// no such tokenId in quest
	var quest model.Quest
	errQuest := global.DB.Model(&model.Quest{}).
		Where("tokenId", claimed.TokenId).
		First(&quest).Error
	if errQuest != nil {
		return errors.New("no such tokenId in quest")
	}
	//
	challenges := model.UserChallenges{
		Address: claimed.Sender.String(),
		QuestID: quest.ID,
		Status:  2,
		Claimed: true,
		ClaimTs: time.Now().Unix(),
	}
	err = global.DB.Model(&model.UserChallenges{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "questId"}},
		UpdateAll: true,
	}).Create(&challenges).Error
	if err != nil {
		return err
	}
	HandleTraverseStatus(hash, 1, "")
	return
}

func AirdropBadge() error {
	client, err := ethclient.Dial(global.CONFIG.BlockChain.Provider)
	if err != nil {
		return errors.New("ethclient dial error")
	}
	paddingList, err := getPendingAirdrop()
	if err != nil {
		return err
	}
	for tokenId, list := range paddingList {
		receivers := receiverNotClaimList(client, tokenId, list)
		hash, err := _airdropBadge(client, tokenId, receivers)
		if err != nil {
			global.LOG.Error("_airdropBadge", zap.Any("error", err))
			continue
		}
		if err := updateAirdropStatus(tokenId, receivers, hash.String()); err != nil {
			global.LOG.Error("updateAirdropStatus", zap.Any("error", err))
		}
	}
	return nil
}

func _airdropBadge(client *ethclient.Client, tokenID int64, receivers []common.Address) (txHash common.Hash, err error) {
	tokenId := big.NewInt(tokenID)

	privateKey, err := crypto.HexToECDSA(global.CONFIG.BlockChain.PrivateKey)
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
			"airdropBadge", tokenId, global.CONFIG.Contract.Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27

	questMinter, err := ABI.NewQuestMinter(common.HexToAddress(global.CONFIG.Contract.QuestMinter), client)
	if err != nil {
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
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
	global.LOG.Info("Airdrop tx sent :", zap.String("hash: ", tx.Hash().Hex()))
	return tx.Hash(), nil
}

func receiverNotClaimList(client *ethclient.Client, tokenId int64, receivers []string) (receiversNotClaim []common.Address) {
	for _, receiver := range receivers {
		badge, err := ABI.NewBadge(common.HexToAddress(global.CONFIG.Contract.Badge), client)
		if err != nil {
			continue
		}
		res, err := badge.BalanceOf(nil, common.HexToAddress(receiver), big.NewInt(tokenId))
		if err != nil {
			continue
		}
		if res.Cmp(big.NewInt(0)) != 0 {
			// already claimed update status
			_ = global.DB.Model(&model.ClaimBadgeTweet{}).
				Where("token_id = ? AND address = ?", tokenId, receiver).
				Update("airdropped", true).Error

			continue
		}
		receiversNotClaim = append(receiversNotClaim, common.HexToAddress(receiver))
	}
	return
}

func getPendingAirdrop() (res map[int64][]string, err error) {
	res = make(map[int64][]string)
	var pending []model.ClaimBadgeTweet
	if err = global.DB.Model(&model.ClaimBadgeTweet{}).Where("airdropped", false).Find(&pending).Error; err != nil {
		return
	}
	for _, v := range pending {
		res[v.TokenId] = append(res[v.TokenId], v.Address)
	}
	return res, nil
}

func updateAirdropStatus(tokenId int64, receivers []common.Address, hash string) error {
	tx := global.DB.Model(&model.ClaimBadgeTweet{}).Begin()
	for _, v := range receivers {
		tx.Where("token_id = ? AND address = ?", tokenId, v.String()).
			Updates(map[string]interface{}{"airdropped": "true", "airdrop_hash": hash, "airdrop_ts": time.Now().Unix()})
	}
	return tx.Commit().Error
}
