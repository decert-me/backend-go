package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"strconv"
	"testing"
)

func TestService_AddQuest(t *testing.T) {
	res, err := s.AddQuest(ADDRESS, request.AddQuestRequest{
		Uri:         "ipfs://Qmd1bCLoEPJ14fuJLgZPGgWh6ravkaV6wQchw71t4y5P2Y",
		Title:       "Title测试123abc@#¥",
		Description: "Description测试123abc@#¥",
	})
	if err != nil {
		t.Error("AddQuest error")
	}
	assert.Equal(t, "0x778a31ae96edf382c47e9b45d99e36482c16cfd956a29aef47877ad819023e457dd69b1dab4ccd371b80a625ed8965105622d8373bd97bd0d0ec02b25972ec6b1b", res)
}

func TestService_GetQuest(t *testing.T) {
	// delete exist
	err := s.dao.DB().Where("token_id", TOKENID).Delete(&model.Quest{}).Error
	assert.Nil(t, err)
	// Start testing
	deleteQuest()
	deleteTransaction()
	s.HashSubmit("", QuestCreatedHash)
	waitForQuestCreated(TOKENID)
	questList, total, err := s.GetQuestList(request.GetQuestListRequest{
		Quest: model.Quest{
			TokenId: TOKENID,
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), total)
	assert.Equal(t, len(questList), 1)
	quest, err := s.GetQuest(strconv.Itoa(TOKENID))
	assert.Nil(t, err)
	metaData := []byte(gjson.Parse("{\"image\": \"ipfs://QmZr8D7Qjurwd8HxzHyqRy7epWsbE5TRUp4yic3qSshGj2\", \"title\": \"Title测试123abc@#¥\", \"version\": 1, \"properties\": {\"url\": \"\", \"answers\": \"PlIZOgVOU28BRxYQEUgWaA==\", \"endTIme\": null, \"requires\": [], \"questions\": [{\"type\": 0, \"score\": 100, \"title\": \"Questions测试123abc@#¥\", \"options\": [\"true\", \"false\"]}, {\"type\": 1, \"score\": 100, \"title\": \"Description2测试123abc@#¥\", \"options\": [\"true\", \"true\", \"false\"]}, {\"type\": 2, \"score\": 100, \"title\": \"Questions3测试123abc@#¥\", \"options\": [\"true\"]}], \"startTime\": \"2023-02-27T02:29:15.240Z\", \"difficulty\": 0, \"estimateTime\": 1800, \"passingScore\": 200}, \"description\": \"Description测试123abc@#¥\"}").Raw)
	extraData := []byte(gjson.Parse("{\"endTs\": 1678327339, \"supply\": 10, \"startTs\": 1}").Raw)
	questExpect := model.Quest{
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

	assert.Equal(t, questExpect, quest)
	assert.Equal(t, questExpect, questList[0].Quest)

	_, err = s.GetQuest("x212")
	assert.Error(t, err, "", "should error when questId not digits")
	// clear
	deleteQuest()
	deleteTransaction()
}
