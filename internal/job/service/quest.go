package service

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
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
	metadata, err := utils.GetDataFromCid(strings.Replace(created.QuestData.Uri, "ipfs://", "", 1))
	if err != nil {
		return
	}
	questData := created.QuestData
	extraData, _ := json.Marshal(model.Extradata{StartTs: questData.StartTs, EndTs: questData.EndTs, Supply: questData.Supply.Uint64()})
	_ = extraData
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
	}
	if err = s.dao.CreateQuest(&quest); err != nil {
		log.Errorv("CreateQuest error", zap.Error(err), zap.Any("quest", quest))
		return
	}
	s.handleTraverseStatus(hash, 1, "")

	return
}
