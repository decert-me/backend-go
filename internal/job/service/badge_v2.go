package service

import (
	ABIV2 "backend-go/abi/v2"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"math/big"
	"strings"
	"time"
)

var badgeAbiV2 abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABIV2.BadgeMinterMetaData.ABI))
	if err != nil {
		panic(err)
	}
	badgeAbiV2 = contractAbi
}

func (s *Service) handleClaimedV2(hash string, vLog *types.Log, chainID int64) (err error) {
	var claimed ABIV2.BadgeClaimed
	if err = badgeAbiV2.UnpackIntoInterface(&claimed, "Claimed", vLog.Data); err != nil || len(vLog.Topics) == 0 {
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
	// TODO: 分数从URL获取
	score := big.NewInt(0)

	challenges := model.UserChallenges{
		Address:   common.HexToAddress(vLog.Topics[2].Hex()).String(),
		TokenId:   tokenId,
		Status:    2,
		UserScore: score.Int64(),
		Claimed:   true,
		ClaimTs:   time.Now().Unix(),
		ChainID:   chainID,
	}
	err = s.dao.CreateChallenges(&challenges)
	if err != nil {
		log.Errorv("CreateChallenges error", zap.Any("challenges", challenges), zap.Error(err))
		return err
	}
	// 如果有空投记录则删除
	s.dao.UpdateAirdroppedError(tokenId, common.HexToAddress(vLog.Topics[2].Hex()).String(), "already claimed")
	s.handleTraverseStatus(hash, 1, "")
	return
}
