package service

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/model"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tidwall/gjson"
	"strings"
)

var badgeAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.BadgeMetaData.ABI))
	if err != nil {
		panic(err)
	}
	badgeAbi = contractAbi
}

func (s *Service) handleQuestInit(trans model.Transaction, vLog *types.Log) (err error) {
	var questInit ABI.BadgeQuestInit
	if err = badgeAbi.UnpackIntoInterface(&questInit, "QuestInit", vLog.Data); err != nil || len(vLog.Topics) == 0 {
		return errors.New("unpack error")
	}
	questId := vLog.Topics[1].Big().Int64()
	quest, err := s.dao.GetQuestByTokenID(questId)
	if err != nil {
		return err
	}
	questNow := questInit.QuestData
	if questNow.Uri == quest.Uri &&
		questNow.StartTs == uint32(gjson.Get(string(quest.MetaData), "startTs").Uint()) &&
		questNow.EndTs == uint32(gjson.Get(string(quest.MetaData), "endTs").Uint()) &&
		questNow.Title == gjson.Get(string(quest.MetaData), "title").String() {
		// 更新状态
		s.dao.SetMultiChainStatus(questId, trans.ChainID, 1)
	}
	s.handleTraverseStatus(trans.Hash, 1, "")
	return
}

func (s *Service) handleQuestUpdated(trans model.Transaction, vLog *types.Log) (err error) {
	var questUpdated ABI.BadgeQuestUpdated
	if err = badgeAbi.UnpackIntoInterface(&questUpdated, "QuestUpdated", vLog.Data); err != nil || len(vLog.Topics) == 0 {
		return errors.New("unpack error")
	}
	questId := vLog.Topics[1].Big().Int64()
	quest, err := s.dao.GetQuestByTokenID(questId)
	if err != nil {
		return err
	}
	questNow := questUpdated.QuestData
	if questNow.Uri == quest.Uri &&
		questNow.StartTs == uint32(gjson.Get(string(quest.MetaData), "startTs").Uint()) &&
		questNow.EndTs == uint32(gjson.Get(string(quest.MetaData), "endTs").Uint()) &&
		questNow.Title == gjson.Get(string(quest.MetaData), "title").String() {
		// 更新状态
		s.dao.SetMultiChainStatus(questId, trans.ChainID, 1)
	}
	s.handleTraverseStatus(trans.Hash, 1, "")
	return
}
