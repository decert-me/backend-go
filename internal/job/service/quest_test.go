package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/job/dao"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"sync"
	"testing"
)

func TestHandleQuestCreated(t *testing.T) {
	deleteQuest()
	s.TaskChain <- taskTx{task: &model.Transaction{Hash: QuestCreatedHash}, txMap: new(sync.Map), countMap: new(sync.Map)}
	waitForQuestCreated(TOKENID)
	quest, err := s.dao.GetQuest(&model.Quest{
		TokenId: TOKENID,
	})
	if err != nil {
		t.Error("GetQuest error")
	}
	if quest.AddTs == 0 {
		t.Error("AddTs error")
	}
	metaData := []byte(gjson.Parse("{\"image\": \"ipfs://QmZr8D7Qjurwd8HxzHyqRy7epWsbE5TRUp4yic3qSshGj2\", \"title\": \"Title测试123abc@#¥\", \"version\": 1, \"properties\": {\"url\": \"\", \"answers\": \"PlIZOgVOU28BRxYQEUgWaA==\", \"endTIme\": null, \"requires\": [], \"questions\": [{\"type\": 0, \"score\": 100, \"title\": \"Questions测试123abc@#¥\", \"options\": [\"true\", \"false\"]}, {\"type\": 1, \"score\": 100, \"title\": \"Description2测试123abc@#¥\", \"options\": [\"true\", \"true\", \"false\"]}, {\"type\": 2, \"score\": 100, \"title\": \"Questions3测试123abc@#¥\", \"options\": [\"true\"]}], \"startTime\": \"2023-02-27T02:29:15.240Z\", \"difficulty\": 0, \"estimateTime\": 1800, \"passingScore\": 200}, \"description\": \"Description测试123abc@#¥\"}").Raw)
	extraData := []byte(gjson.Parse("{\"endTs\": 1886306695, \"supply\": 10, \"startTs\": 1}").Raw)
	questAssert := model.Quest{
		ID:           quest.ID,
		Title:        "以太坊简介",
		Label:        "",
		Disabled:     false,
		Description:  "Description测试123abc@#¥",
		Dependencies: nil,
		IsDraft:      false,
		AddTs:        quest.AddTs,
		TokenId:      TOKENID,
		Type:         0,
		Difficulty:   0,
		EstimateTime: 0,
		Creator:      ADDRESS,
		MetaData:     metaData,
		ExtraData:    extraData,
		Uri:          "ipfs://QmU1TsKEEh1dMEkArp4tkKqf9qU1t9CCCvG6B1py879rhb",
	}
	assert.Equal(t, questAssert, quest)

	// clean
	deleteQuest()
}

func TestBlockChain_handleQuestCreated(t *testing.T) {
	err := s.handleQuestCreated("", &types.Log{})
	assert.Error(t, err, "should return error when error Log")
}

func TestQuestServiceCrash(t *testing.T) {
	s.dao.Close() // Service Crash
	// Start testing
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa"}, txMap: new(sync.Map), countMap: new(sync.Map)})
	// restart
	s.dao = dao.New(c)
}
