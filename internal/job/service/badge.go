package service

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tidwall/gjson"
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
	metadata, err := s.GetDataFromCid(strings.Replace(uri.Value, "ipfs://", "", 1))
	if err != nil {
		return
	}
	var questDataDetail string
	if gjson.Get(metadata, "version").Float() == 1.1 {
		questDataDetail, err = s.GetDataFromCid(strings.Replace(gjson.Get(metadata, "attributes.challenge_ipfs_url").String(), "ipfs://", "", 1))
		if err != nil {
			return
		}
	}

	challengeUrl := gjson.Get(metadata, "attributes.challenge_url").String()
	var uuid string
	if len(strings.Split(challengeUrl, "/quests/")) >= 2 {
		uuid = strings.Split(challengeUrl, "/quests/")[1]
	}

	quest := model.Quest{
		UUID:      uuid,
		TokenId:   tokenId,
		Uri:       uri.Value,
		MetaData:  []byte(metadata),
		QuestData: []byte(questDataDetail),
	}
	err = s.dao.UpdateQuest(&quest)
	if err != nil {
		log.Errorv("UpdateQuest error", zap.Any("quest", quest), zap.Error(err))
		return err
	}
	s.handleTraverseStatus(hash, 1, "")
	return
}
