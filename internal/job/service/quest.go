package service

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"encoding/json"
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
	tr, err := s.dao.QueryTransactionByHash(hash)
	if err != nil {
		return err
	}
	questData := created.QuestData
	extraData, _ := json.Marshal(model.Extradata{StartTs: questData.StartTs,
		EndTs:   questData.EndTs,
		Title:   questData.Title,
		Creator: common.HexToAddress(vLog.Topics[1].Hex()),
	})
	// TODO: 多链状态
	for _, v := range s.c.Contract.MultiChain {
		v.
	}
	multiChainStatus
	quest := model.Quest{
		Title:       questData.Title,
		Description: gjson.Get(metadata, "description").String(),
		TokenId:     vLog.Topics[2].Big().Int64(),
		Uri:         questData.Uri,
		Type:        0, // TODO
		Creator:     common.HexToAddress(vLog.Topics[1].Hex()).String(),
		MetaData:    []byte(metadata),
		ExtraData:   extraData,
		IsDraft:     false, // 当前发布不审核
		Recommend:   gjson.Get(tr.Params.String(), "recommend").Raw,
	}
	if err = s.dao.CreateQuest(&quest); err != nil {
		log.Errorv("CreateQuest error", zap.Error(err), zap.Any("quest", quest))
		return
	}
	s.handleTraverseStatus(hash, 1, "")

	return
}

func (s *Service) handleModifyQuest(hash string, vLog *types.Log) (err error) {
	var modify ABI.QuestQuestModify
	if err = questAbi.UnpackIntoInterface(&modify, "QuestModify", vLog.Data); err != nil {
		return
	}
	metadata, err := s.GetDataFromCid(strings.Replace(modify.QuestData.Uri, "ipfs://", "", 1))
	if err != nil {
		return
	}
	tr, err := s.dao.QueryTransactionByHash(hash)
	if err != nil {
		return err
	}
	questData := modify.QuestData
	extraData, _ := json.Marshal(model.Extradata{StartTs: questData.StartTs,
		EndTs:   questData.EndTs,
		Title:   questData.Title,
		Creator: common.HexToAddress(vLog.Topics[1].Hex()),
	})
	quest := model.Quest{
		Title:       questData.Title,
		Description: gjson.Get(metadata, "description").String(),
		TokenId:     vLog.Topics[2].Big().Int64(),
		Uri:         questData.Uri,
		Type:        0, // TODO
		Creator:     common.HexToAddress(vLog.Topics[1].Hex()).String(),
		MetaData:    []byte(metadata),
		ExtraData:   extraData,
		IsDraft:     false, // 当前发布不审核
		Recommend:   gjson.Get(tr.Params.String(), "recommend").Raw,
	}
	if err = s.dao.UpdateQuest(&quest); err != nil {
		log.Errorv("UpdateQuest error", zap.Error(err), zap.Any("quest", quest))
		return
	}

	s.handleTraverseStatus(hash, 1, "")

	return
}
