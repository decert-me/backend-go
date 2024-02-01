package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinzhu/copier"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"math/big"
)

func (s *Service) GetQuestList(searchInfo request.GetQuestListRequest) (res []response.GetQuestListRes, total int64, err error) {
	res, total, err = s.dao.GetQuestList(&searchInfo)
	return
}

func (s *Service) GetUserQuestList(searchInfo request.GetUserQuestListRequest) (res []response.GetUserQuestListRes, total int64, err error) {
	res, total, err = s.dao.GetUserQuestList(&searchInfo)
	return
}

func (s *Service) GetUserQuestListWithClaimed(searchInfo request.GetUserQuestListRequest) (res []response.QuestWithClaimed, total int64, err error) {
	res, total, err = s.dao.GetUserQuestListWithClaimed(&searchInfo)
	return
}

func (s *Service) GetQuest(language, id string, address, original string) (quest response.GetQuestRes, err error) {
	if utils.IsUUID(id) {
		if address == "" {
			quest, err = s.dao.GetQuestByUUIDLang(language, id)
			return
		}
		quest, err = s.dao.GetQuestWithClaimStatusByUUID(language, id, address)
	} else {
		if address == "" {
			quest, err = s.dao.GetQuestByTokenIDWithLang(language, id)
			return
		}
		if original == "true" {
			quest, err = s.dao.GetQuestWithClaimStatusByTokenID(id, address)
			return
		} else {
			quest, err = s.dao.GetQuestWithClaimStatusByTokenIDWithLang(language, id, address)
		}
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
	tokenId, set := big.NewInt(0).SetString(modify.TokenId, 10)
	if !set {
		return res, errors.New("TokenIDInvalid")
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "uint32", "uint32", "uint192", "string", "string", "address", "address"},
		// values
		[]interface{}{
			tokenId, modify.StartTs, modify.EndTs, modify.Supply, modify.Title, modify.Uri, s.c.Contract.QuestMinter, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func (s *Service) GetQuestChallengeUser(id string) (res response.GetQuestChallengeUserRes, err error) {
	if utils.IsUUID(id) {
		res, err = s.dao.GetQuestChallengeUserByUUID(id)
	} else {
		res, err = s.dao.GetQuestChallengeUserByTokenID(id)
	}
	return
}

func (s *Service) UpdateRecommend(address string, modify request.UpdateRecommendRequest) (err error) {
	// 获取Quest信息
	quest, err := s.dao.GetQuestByTokenID(modify.TokenId)
	if err != nil {
		return errors.New("UnexpectedError")
	}
	if quest.Creator != common.HexToAddress(address).String() {
		return errors.New("UnauthorizedAccess")
	}
	// 修改Quest
	err = s.dao.UpdateQuest(&model.Quest{
		TokenId:   modify.TokenId,
		Recommend: modify.Recommend,
	})
	if err != nil {
		return errors.New("OperationFailed")
	}
	return nil
}

// GetQuestFlashRank 获取闪电榜
func (s *Service) GetQuestFlashRank(address string, id string) (res response.GetQuestFlashListRes, err error) {
	if utils.IsUUID(id) {
		res, err = s.dao.GetQuestFlashRankByUUID(address, id)
	} else {
		res, err = s.dao.GetQuestFlashRankByTokenID(address, id)
	}
	return
}

// GetQuestHighRank 获取高分榜
func (s *Service) GetQuestHighRank(address string, id string) (res response.GetQuestHighScoreListRes, err error) {
	if utils.IsUUID(id) {
		res, err = s.dao.GetQuestHighRankByUUID(address, id)
	} else {
		res, err = s.dao.GetQuestHighRankByTokenID(address, id)
	}
	return
}

// GetQuestHolderRank 获取持有榜
func (s *Service) GetQuestHolderRank(address, id string, page int, pageSize int) (res []response.GetQuestHolderListRes, total int64, err error) {
	if utils.IsUUID(id) {
		res, total, err = s.dao.GetQuestHolderRankByUUID(address, id, page, pageSize)
	} else {
		res, total, err = s.dao.GetQuestHolderRankByTokenID(address, id, page, pageSize)
	}
	return
}

// GetQuestUserScore 获取用户分数
func (s *Service) GetQuestUserScore(id, address string) (res response.GetQuestUserScoreRes, err error) {
	if utils.IsUUID(id) {
		data, err := s.dao.GetQuestWithClaimStatusByUUID("", id, address)
		if err != nil {
			return res, err
		}
		copier.Copy(&res, &data)
	} else {
		data, err := s.dao.GetQuestWithClaimStatusByTokenIDWithLang("", id, address)
		if err != nil {
			return res, err
		}
		copier.Copy(&res, &data)
	}
	res.Address = address
	return
}
