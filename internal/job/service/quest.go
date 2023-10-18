package service

import (
	ABI "backend-go/abi"
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

var questAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.QuestMetaData.ABI))
	if err != nil {
		panic(err)
	}
	questAbi = contractAbi
}

func (s *Service) handleQuestCreated(hash string, vLog *types.Log) (err error) {
	var created ABI.QuestQuestCreated
	if err = questAbi.UnpackIntoInterface(&created, "QuestCreated", vLog.Data); err != nil {
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
	extraData, _ := json.Marshal(model.Extradata{StartTs: questData.StartTs, EndTs: questData.EndTs, Supply: questData.Supply.Uint64()})

	challengeUrl := gjson.Get(metadata, "attributes.challenge_url").String()
	var uuid string
	if len(strings.Split(challengeUrl, "/quests/")) >= 2 {
		uuid = strings.Split(challengeUrl, "/quests/")[1]
	}

	quest := model.Quest{
		UUID:        uuid,
		Title:       questData.Title,
		Description: gjson.Get(metadata, "description").String(),
		TokenId:     vLog.Topics[2].Big().Int64(),
		Uri:         questData.Uri,
		Type:        0, // TODO
		Creator:     common.HexToAddress(vLog.Topics[1].Hex()).String(),
		MetaData:    []byte(metadata),
		ExtraData:   extraData,
		QuestData:   []byte(questDataDetail),
		IsDraft:     false, // 当前发布不审核
		Recommend:   gjson.Get(tr.Params.String(), "recommend").String(),
	}
	// 区分合辑和Quest
	if collectionID == 0 {
		if err = s.dao.CreateQuest(&quest); err != nil {
			log.Errorv("CreateQuest error", zap.Error(err), zap.Any("quest", quest))
			return
		}
	} else {
		if err = s.dao.UpdateCollection(collectionID, quest); err != nil {
			log.Errorv("UpdateCollection error", zap.Error(err), zap.Any("quest", quest))
			return
		}
	}

	s.handleTraverseStatus(hash, 1, "")

	return
}

func (s *Service) handleModifyQuest(hash string, resJson []byte) (err error) {
	tr, err := s.dao.QueryTransactionByHash(hash)
	if err != nil {
		return err
	}
	var questData ABI.IQuestQuestData
	err = json.Unmarshal([]byte(gjson.Get(string(resJson), "questData").String()), &questData)
	if err != nil {
		fmt.Println(err)
		return
	}
	metadata, err := s.GetDataFromCid(strings.Replace(questData.Uri, "ipfs://", "", 1))
	if err != nil {
		return
	}
	var questDataDetail string
	version := gjson.Get(metadata, "version").Float()
	if version == 1.1 || version == 1.2 {
		questDataDetail, err = s.GetDataFromCid(strings.Replace(gjson.Get(metadata, "attributes.challenge_ipfs_url").String(), "ipfs://", "", 1))
		if err != nil {
			return
		}
	}
	extraData, _ := json.Marshal(model.Extradata{StartTs: questData.StartTs, EndTs: questData.EndTs, Supply: questData.Supply.Uint64()})
	quest := model.Quest{
		Title:       questData.Title,
		Description: gjson.Get(metadata, "description").String(),
		TokenId:     gjson.Get(string(resJson), "tokenId").Int(),
		Uri:         questData.Uri,
		Type:        0, // TODO
		MetaData:    []byte(metadata),
		ExtraData:   extraData,
		Recommend:   gjson.Get(tr.Params.String(), "recommend").Raw,
		QuestData:   []byte(questDataDetail),
	}
	if err = s.dao.UpdateQuest(&quest); err != nil {
		log.Errorv("UpdateQuest error", zap.Error(err), zap.Any("quest", quest))
		return
	}

	s.handleTraverseStatus(hash, 1, "")

	return
}
