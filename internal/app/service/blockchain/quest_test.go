package blockchain

import (
	"backend-go/internal/app/model"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
)

func TestHandleQuestCreated(t *testing.T) {
	deleteQuest()
	b.TaskChain <- model.Transaction{Hash: "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa"}
	waitForQuestCreated(10003)
	quest, err := b.dao.GetQuest(&model.Quest{
		TokenId: 10003,
	})
	if err != nil {
		t.Error("GetQuest error")
	}
	if quest.AddTs == 0 {
		t.Error("AddTs error")
	}
	metaData := []byte(gjson.Parse("{\"image\": \"ipfs://QmZr8D7Qjurwd8HxzHyqRy7epWsbE5TRUp4yic3qSshGj2\", \"title\": \"Title测试123abc@#¥\", \"version\": 1, \"properties\": {\"url\": \"\", \"answers\": \"C0oUE34dfRxoFSVANlNFKw==\", \"endTIme\": null, \"requires\": [], \"questions\": [{\"type\": 0, \"score\": 100, \"title\": \"Questions测试123abc@#¥\", \"options\": [\"true\", \"false\"]}, {\"type\": 1, \"score\": 100, \"title\": \"Description2测试123abc@#¥\", \"options\": [\"true\", \"true\", \"false\"]}, {\"type\": 2, \"score\": 100, \"title\": \"Questions3测试123abc@#¥\", \"options\": [\"true\"]}], \"startTime\": \"2023-02-27T02:29:15.240Z\", \"difficulty\": 0, \"estimateTime\": 1800, \"passingScore\": 200}, \"description\": \"Description测试123abc@#¥\"}").Raw)
	extraData := []byte(gjson.Parse("{\"endTs\": 0, \"supply\": 0, \"startTs\": 0}").Raw)
	questAssert := model.Quest{
		ID:           quest.ID,
		Title:        "Title测试123abc@#¥",
		Label:        "",
		Disabled:     false,
		Description:  "Description测试123abc@#¥",
		Dependencies: nil,
		IsDraft:      false,
		AddTs:        quest.AddTs,
		TokenId:      10003,
		Type:         0,
		Difficulty:   0,
		EstimateTime: 0,
		Creator:      "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
		MetaData:     metaData,
		ExtraData:    extraData,
		Uri:          "ipfs://Qmd1bCLoEPJ14fuJLgZPGgWh6ravkaV6wQchw71t4y5P2Y",
	}
	assert.Equal(t, questAssert, quest)
}

func TestBlockChain_handleQuestCreated(t *testing.T) {
	err := b.handleQuestCreated("", &types.Log{})
	fmt.Println(err)
	assert.Error(t, err, "should return error when error Log")
}
