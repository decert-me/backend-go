package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (s *Service) PermitClaimBadge(address string, req request.PermitClaimBadgeReq) (res string, err error) {
	// 校验分数正确性
	quest, err := s.dao.GetQuestByTokenID(req.TokenId)
	if err != nil {
		return res, errors.New("TokenIDInvalid")
	}
	// 校验题目
	if req.StandardAnswer != gjson.Get(string(quest.QuestData), "answers").String() {
		return res, errors.New("TokenIDInvalid")
	}
	pass, err := s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, req.Score, &quest)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return res, errors.New("UnexpectedError")
	}
	if !pass {
		return res, errors.New("AnswerIncorrect")
	}
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		[]string{"string", "uint256", "uint256", "address", "address"},
		[]interface{}{
			"claim", big.NewInt(req.TokenId), big.NewInt(req.Score), s.c.Contract.Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func (s *Service) SubmitClaimTweet(address string, req request.SubmitClaimTweetReq) (err error) {
	// 检查tokenId是否存在以及可用
	valid, err := s.dao.ValidTokenId(req.TokenId)
	if err != nil {
		log.Errorv("ValidTokenId error", zap.Int64("TokenId", req.TokenId), zap.Error(err))
		return errors.New("TokenIDInvalid")
	}
	if !valid {
		return errors.New("TokenIDInvalid")
	}
	// 校验分数正确性
	quest, err := s.dao.GetQuestByTokenID(req.TokenId)
	if err != nil {
		return errors.New("TokenIDInvalid")
	}
	pass, err := s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, req.Score, &quest)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	if !pass {
		return errors.New("AnswerIncorrect")
	}
	// 获取推文ID
	tweetId := utils.GetTweetIdFromURL(req.TweetUrl)
	if tweetId == "" {
		return errors.New("BrokenLink")
	}
	// 检查是否重复使用
	used, err := s.dao.HasTweet(tweetId)
	if err != nil {
		log.Errorv("HasTweet error", zap.Int64("TokenId", req.TokenId), zap.Error(err))
		return errors.New("UnexpectedError")
	}
	if used {
		return errors.New("TweetRepeated")
	}
	// 保存到数据库
	exists, err := s.dao.CreateClaimBadgeTweet(&model.ClaimBadgeTweet{
		Address: address,
		TokenId: req.TokenId,
		Score:   req.Score,
		Url:     req.TweetUrl,
		TweetId: tweetId,
		AddTs:   time.Now().Unix(),
	})
	if err != nil {
		log.Errorv("CreateClaimBadgeTweet error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	if exists {
		return errors.New("AlreadyHoldsBadge")
	}
	return nil
}

func (s *Service) UpdateBadgeURI(address string, badgeURI request.UpdateBadgeURIRequest) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "string", "address", "address"},
		// values
		[]interface{}{
			big.NewInt(badgeURI.TokenId), badgeURI.Uri, s.c.Contract.Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}
