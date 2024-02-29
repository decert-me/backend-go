package service

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

func (s *Service) handleClaimed(client *ethclient.Client, hash string, vLog *types.Log) (err error) {
	var claimed ABI.QuestMinterClaimed
	if err = questMinterAbi.UnpackIntoInterface(&claimed, "Claimed", vLog.Data); err != nil || len(vLog.Topics) == 0 {
		return errors.New("unpack error")
	}
	tokenId := vLog.Topics[1].Big().String()
	// no such tokenId in quest
	exist, err := s.dao.HasTokenId(tokenId)
	if err != nil {
		log.Errorv("HasTokenId error", zap.String("tokenId", tokenId), zap.Error(err))
		return
	}
	if !exist {
		log.Errorv("no such tokenId in quest", zap.String("tokenId", tokenId))
		return
	}
	// 获取用户分数
	badge, err := ABI.NewBadge(common.HexToAddress(s.c.Contract.Badge), client)
	if err != nil {
		return
	}
	tokenIDInt, _ := new(big.Int).SetString(tokenId, 10)

	score, err := badge.Scores(nil, tokenIDInt, common.HexToAddress(vLog.Topics[2].Hex()))
	if err != nil {
		return
	}
	challenges := model.UserChallenges{
		Address:   common.HexToAddress(vLog.Topics[2].Hex()).String(),
		TokenId:   tokenId,
		Status:    2,
		UserScore: score.Int64(),
		Claimed:   true,
		ClaimTs:   time.Now().Unix(),
	}
	err = s.dao.CreateChallenges(&challenges)
	if err != nil {
		log.Errorv("CreateChallenges error", zap.Any("challenges", challenges), zap.Error(err))
		return err
	}
	s.handleTraverseStatus(hash, 1, "")
	return
}
