package service

import (
	ABIV2 "backend-go/abi/v2"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"strings"
)

var questAbiV2 abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABIV2.QuestMetaData.ABI))
	if err != nil {
		panic(err)
	}
	questAbiV2 = contractAbi
}

func (s *Service) handleQuestCreatedV2(hash string, vLog *types.Log) (err error) {
	var created ABIV2.QuestQuestCreated
	if err = questAbiV2.UnpackIntoInterface(&created, "QuestCreated", vLog.Data); err != nil {
		return
	}
	metadata, err := s.GetDataFromCid(strings.Replace(created.QuestData.Uri, "ipfs://", "", 1))
	if err != nil {
		return
	}
	// 获取数据库信息
	tr, err := s.dao.QueryTransactionByHash(hash)
	if err != nil {
		return err
	}
	// 获取合辑ID
	collectionID := gjson.Get(tr.Params.String(), "collection_id").Int()
	var questDataDetail string
	version := gjson.Get(metadata, "version").Float()
	if (version == 1.1 || version == 1.2) && collectionID == 0 {
		questDataDetail, err = s.GetDataFromCid(strings.Replace(gjson.Get(metadata, "attributes.challenge_ipfs_url").String(), "ipfs://", "", 1))
		if err != nil {
			return
		}
	}

	questData := created.QuestData
	extraData, _ := json.Marshal(model.Extradata{StartTs: questData.StartTs, EndTs: questData.EndTs})

	challengeUrl := gjson.Get(metadata, "attributes.challenge_url").String()
	var uuid string
	if collectionID == 0 {
		if len(strings.Split(challengeUrl, "/quests/")) >= 2 {
			uuid = strings.Split(challengeUrl, "/quests/")[1]
		}
	} else {
		if len(strings.Split(challengeUrl, "/collection/")) >= 2 {
			uuid = strings.Split(challengeUrl, "/collection/")[1]
		}
	}

	quest := model.Quest{
		UUID:        uuid,
		Title:       questData.Title,
		Description: gjson.Get(metadata, "description").String(),
		TokenId:     vLog.Topics[2].Big().String(),
		Uri:         questData.Uri,
		Type:        0, // TODO
		Creator:     common.HexToAddress(vLog.Topics[1].Hex()).String(),
		MetaData:    []byte(metadata),
		ExtraData:   extraData,
		QuestData:   []byte(questDataDetail),
		IsDraft:     false, // 当前发布不审核
		Recommend:   gjson.Get(tr.Params.String(), "recommend").String(),
		Version:     "2",
		ChainID:     tr.ChainID,
	}

	// 区分合辑和Quest
	if collectionID == 0 {
		if err = s.dao.CreateQuest(&quest); err != nil {
			log.Errorv("CreateQuest error", zap.Error(err), zap.Any("quest", quest))
			return
		}
	} else {
		if err = s.dao.UpdateCollectionOnce(collectionID, quest); err != nil {
			log.Errorv("UpdateCollectionOnce error", zap.Error(err), zap.Any("quest", quest))
			return
		}
	}

	s.handleTraverseStatus(hash, 1, "")

	return
}

func (s *Service) handleModifyQuestV2(hash string, vLog *types.Log) (err error) {
	var modified ABIV2.QuestQuestModified
	if err = questAbiV2.UnpackIntoInterface(&modified, "QuestModified", vLog.Data); err != nil {
		fmt.Println(err)
		return
	}
	metadata, err := s.GetDataFromCid(strings.Replace(modified.QuestData.Uri, "ipfs://", "", 1))
	if err != nil {
		return
	}
	tr, err := s.dao.QueryTransactionByHash(hash)
	if err != nil {
		return err
	}
	var questDataDetail string
	version := gjson.Get(metadata, "version").Float()
	if version == 1.1 || version == 1.2 {
		questDataDetail, err = s.GetDataFromCid(strings.Replace(gjson.Get(metadata, "attributes.challenge_ipfs_url").String(), "ipfs://", "", 1))
		if err != nil {
			return
		}
	}
	extraData, _ := json.Marshal(model.Extradata{StartTs: modified.QuestData.StartTs, EndTs: modified.QuestData.EndTs})
	quest := model.Quest{
		Title:       modified.QuestData.Title,
		Description: gjson.Get(metadata, "description").String(),
		TokenId:     vLog.Topics[2].Big().String(),
		Uri:         modified.QuestData.Uri,
		Type:        0, // TODO
		MetaData:    []byte(metadata),
		ExtraData:   extraData,
		Recommend:   gjson.Get(tr.Params.String(), "recommend").String(),
		QuestData:   []byte(questDataDetail),
		Version:     "2",
	}
	if err = s.dao.UpdateQuest(&quest); err != nil {
		log.Errorv("UpdateQuest error", zap.Error(err), zap.Any("quest", quest))
		return
	}
	// 清除挑战记录
	if err = s.dao.DeleteUserChallengeLogByTokenId(quest.TokenId); err != nil {
		log.Errorv("DeleteUserChallengeLogByTokenId error", zap.Error(err), zap.Any("quest", quest))
		return
	}
	s.handleTraverseStatus(hash, 1, "")

	return
}
