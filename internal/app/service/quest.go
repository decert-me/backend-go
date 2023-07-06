package service

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"math/big"
	"strconv"
)

func (s *Service) GetQuestList(searchInfo request.GetQuestListRequest) (res []response.GetQuestListRes, total int64, err error) {
	res, total, err = s.dao.GetQuestList(&searchInfo)
	return
}

func (s *Service) GetUserQuestList(searchInfo request.GetUserQuestListRequest) (res []response.GetUserQuestListRes, total int64, err error) {
	res, total, err = s.dao.GetUserQuestList(&searchInfo)
	return
}

func (s *Service) GetUserQuestListWithClaimed(searchInfo request.GetUserQuestListRequest) (res []response.GetUserQuestListRes, total int64, err error) {
	res, total, err = s.dao.GetUserQuestListWithClaimed(&searchInfo)
	return
}

func (s *Service) GetQuest(id string, address string) (quest response.GetQuestRes, err error) {
	tokenId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		if address == "" {
			quest.Quest, err = s.dao.GetQuestByUUID(id)
			return
		}
		quest, err = s.dao.GetQuestWithClaimStatusByUUID(id, address)
	} else {
		if address == "" {
			quest.Quest, err = s.dao.GetQuestByTokenID(tokenId)
			return
		}
		quest, err = s.dao.GetQuestWithClaimStatusByTokenID(tokenId, address)
	}

	return
}

func (s *Service) AddQuest(address string, add request.AddQuestRequest) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	//supply, _ := new(big.Int).SetString(add.Supply, 10)
	fmt.Println(add.Supply)
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint32", "uint32", "uint192", "string", "string", "address", "address"},
		// values
		[]interface{}{
			add.StartTs, add.EndTs, add.Supply, add.Title, add.Uri, s.c.Contract.QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func (s *Service) UpdateQuest(address string, modify request.UpdateQuestRequest) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "uint32", "uint32", "uint192", "string", "string", "address", "address"},
		// values
		[]interface{}{
			big.NewInt(modify.TokenId), modify.StartTs, modify.EndTs, modify.Supply, modify.Title, modify.Uri, s.c.Contract.QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func (s *Service) GetQuestChallengeUser(id string) (res response.GetQuestChallengeUserRes, err error) {
	tokenId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		res, err = s.dao.GetQuestChallengeUserByUUID(id)
	} else {
		res, err = s.dao.GetQuestChallengeUserByTokenID(tokenId)
	}

	return
}
