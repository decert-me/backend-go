package service

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
	"strings"
)

func (s *Service) handleURI(hash string, vLog *types.Log) (err error) {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.IERC1155MetaData.ABI))
	if err != nil {
		return
	}
	var uri ABI.BadgeURI
	if err = contractAbi.UnpackIntoInterface(&uri, "URI", vLog.Data); err != nil || len(vLog.Topics) == 0 {
		return errors.New("unpack error")
	}
	tokenId := vLog.Topics[1].Big().Int64()
	// no such tokenId in quest
	exist, err := s.dao.HasTokenId(tokenId)
	if err != nil {
		log.Errorv("HasTokenId error", zap.Int64("tokenId", tokenId), zap.Error(err))
		return
	}
	if !exist {
		log.Errorv("no such tokenId in quest", zap.Int64("tokenId", tokenId))
		return errors.New("no such tokenId in quest")
	}
	metadata, err := utils.GetDataFromCid(strings.Replace(uri.Value, "ipfs://", "", 1))
	if err != nil {
		return
	}
	quest := model.Quest{
		TokenId:  tokenId,
		Uri:      uri.Value,
		MetaData: []byte(metadata),
	}
	err = s.dao.UpdateQuest(&quest)
	if err != nil {
		log.Errorv("UpdateQuest error", zap.Any("quest", quest), zap.Error(err))
		return err
	}
	s.handleTraverseStatus(hash, 1, "")
	return
}
