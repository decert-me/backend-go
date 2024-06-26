package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"github.com/spf13/cast"
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
	if req.Uri != "" && req.Uri != quest.Uri {
		return res, errors.New("QuestUpdate")
	}
	_, _, _, pass, err := s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, address, req.Score, &quest, true)
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
	tokenId, set := big.NewInt(0).SetString(req.TokenId, 10)
	if !set {
		return res, errors.New("TokenIDInvalid")
	}
	hash := solsha3.SoliditySHA3(
		[]string{"string", "uint256", "uint256", "address", "address"},
		[]interface{}{
			"claim", tokenId, big.NewInt(req.Score), s.c.Contract.Badge, address,
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
		log.Errorv("ValidTokenId error", zap.String("TokenId", req.TokenId), zap.Error(err))
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
	// 校验题目
	if req.Uri != "" && req.Uri != quest.Uri {
		return errors.New("QuestUpdate")
	}
	_, _, _, pass, err := s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, address, req.Score, &quest, true)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	if !pass {
		return errors.New("AnswerIncorrect")
	}
	if req.TweetUrl == "" {
		// 保存到数据库
		exists, err := s.dao.CreateClaimBadgeTweet(&model.ClaimBadgeTweet{
			Address: address,
			TokenId: req.TokenId,
			Score:   req.Score,
			Url:     req.TweetUrl,
			AddTs:   time.Now().Unix(),
			Type:    1,
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
	// 获取推文ID
	tweetId := utils.GetTweetIdFromURL(req.TweetUrl)
	if tweetId == "" {
		return errors.New("BrokenLink")
	}
	// 检查是否重复使用
	used, err := s.dao.HasTweet(tweetId)
	if err != nil {
		log.Errorv("HasTweet error", zap.String("TokenId", req.TokenId), zap.Error(err))
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
	tokenId, set := big.NewInt(0).SetString(badgeURI.TokenId, 10)
	if !set {
		return res, errors.New("TokenIDInvalid")
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "string", "address", "address"},
		// values
		[]interface{}{
			tokenId, badgeURI.Uri, s.c.Contract.Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func (s *Service) SubmitClaimShare(address string, req request.SubmitClaimShareReq) (res string, err error) {
	// 校验是否绑定社交账号
	data, err := s.dao.HasBindSocialAccount(address)
	var isBinding bool
	for _, v := range data {
		if v == true {
			isBinding = true
			break
		}
	}
	if !isBinding {
		return res, errors.New("NoBindingDetected")
	}
	// 校验是否已经空投
	if s.dao.HasAirdrop(address, req.TokenId) {
		return res, errors.New("AlreadyAirdrop")
	}
	// 校验分数正确性
	quest, err := s.dao.GetQuestByTokenID(req.TokenId)
	if err != nil {
		return res, errors.New("TokenIDInvalid")
	}
	// 校验题目
	if req.Uri != "" && req.Uri != quest.Uri {
		return res, errors.New("QuestUpdate")
	}
	_, _, _, pass, err := s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, address, req.Score, &quest, true)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return res, errors.New("UnexpectedError")
	}
	if !pass {
		return res, errors.New("AnswerIncorrect")
	}
	var app string
	// 判断地址
	if utils.IsValidAddress(address) {
		app = "decert"
	} else {
		app = "decert_solana"
	}
	// 生成分享码
	paramsMap := map[string]interface{}{
		"app": app,
		"params": map[string]interface{}{
			"receiver": address,
			"tokenId":  req.TokenId,
			"score":    req.Score,
			"uri":      quest.Uri,
		},
	}
	// 将Map转换为JSON格式的字节数组
	paramsData, err := json.Marshal(paramsMap)
	if err != nil {
		log.Errorv("JSON encoding error:", zap.Error(err))
		return
	}
	res, err = s.GenerateShare(request.GenerateShareRequest{Params: string(paramsData)})
	if err != nil {
		log.Errorv("GenerateShare error:", zap.Error(err))
		return
	}
	// 保存记录
	if err = s.dao.CreateUserChallengeClaim(&model.UserChallengeClaim{
		Address: address,
		TokenId: cast.ToString(req.TokenId),
	}); err != nil {
		log.Errorv("CreateUserChallengeClaim error", zap.Error(err))
		return
	}
	return res, err
}

func (s *Service) HasClaimed(address string, tokenId string) (res uint8, err error) {
	// 校验是否已经空投
	if s.dao.HasAirdrop(address, tokenId) {
		return 2, nil
	}
	return s.dao.HasClaimed(address, tokenId)
}
