package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"strconv"
)

func (s *Service) GetQuestList(searchInfo request.GetQuestListRequest) (res []response.GetQuestListRes, total int64, err error) {
	var questList []model.Quest
	if questList, total, err = s.dao.GetQuestList(&searchInfo); err != nil {
		return
	}
	for _, v := range questList {
		res = append(res, response.GetQuestListRes{Quest: v})
	}
	return
}

func (s *Service) GetQuest(id string) (quest model.Quest, err error) {
	tokenId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	quest, err = s.dao.GetQuest(&model.Quest{TokenId: tokenId})
	return
}

func (s *Service) AddQuest(address string, add request.AddQuestRequest) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.PrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint32", "uint32", "uint192", "string", "string", "address", "address"},
		// values
		[]interface{}{
			0, 0, 0, add.Title, add.Uri, s.c.Contract.QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}
