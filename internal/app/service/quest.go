package service

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func GetQuestList(searchInfo request.GetQuestListRequest) (questList []response.GetQuestListRes, err error) {
	var quest []model.Quest
	db := global.DB.Model(&model.Quest{}).Where(&searchInfo.Quest)
	err = db.Order("id desc").Find(&quest).Error
	if err != nil {
		return questList, err
	}
	for _, v := range quest {
		questList = append(questList, response.GetQuestListRes{Quest: v})
	}
	return
}

func GetQuest(id string) (questList response.GetQuestListRes, err error) {
	db := global.DB.Model(&model.Quest{}).Where("tokenId", id)
	err = db.Order("id desc").First(&questList.Quest).Error
	return
}

func AddQuest(address string, add request.AddQuestRequest) (res string, err error) {
	verify := model.Signature{MessageId: "questSubmit", Address: address, Uri: add.Uri}
	verifyByte, err := json.Marshal(verify)
	if err != nil || !utils.VerifySignature(address, add.Signature, verifyByte) {
		fmt.Println("Error signing request")
		return
	}
	privateKey, err := crypto.HexToECDSA(global.CONFIG.BlockChain.PrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint32", "uint32", "uint192", "string", "string", "address", "address"},
		// values
		[]interface{}{
			0, 0, 0, add.Title, add.Uri, global.CONFIG.Contract.QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}
