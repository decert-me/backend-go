package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (s *Service) PermitClaimBadge(address string, req request.PermitClaimBadgeReq) (res string, err error) {
	// 校验分数正确性
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenId})
	if err != nil {
		return res, errors.New("tokenID invalid")
	}
	score, pass, err := utils.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, quest.Uri)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return res, errors.New("something error")
	}
	if !pass || score < req.Score {
		return res, errors.New("answer error")
	}
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.PrivateKey)
	if err != nil {
		return
	}
	hash := solsha3.SoliditySHA3(
		[]string{"uint256", "uint256", "address", "address"},
		[]interface{}{
			big.NewInt(req.TokenId), big.NewInt(req.Score), s.c.Contract.Badge, address,
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
	fmt.Println(valid)
	fmt.Println(err)
	if err != nil {
		log.Errorv("ValidTokenId error", zap.Int64("TokenId", req.TokenId), zap.Error(err))
		return errors.New("ValidTokenId error")
	}
	if !valid {
		return errors.New("invalid quest")
	}
	// TODO: 检查用户是否已通过挑战，或已领取SBT
	// 校验分数正确性
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenId})
	if err != nil {
		return errors.New("tokenID invalid")
	}
	score, pass, err := utils.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, quest.Uri)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return errors.New("something error")
	}
	if !pass || score < req.Score {
		return errors.New("answer error")
	}
	// 获取推文ID
	tweetId := utils.GetTweetIdFromURL(req.TweetUrl)
	if tweetId == "" {
		return errors.New("cannot find tweet id")
	}
	// 检查是否重复使用
	used, err := s.dao.HasTweet(tweetId)
	if err != nil {
		log.Errorv("HasTweet error", zap.Int64("TokenId", req.TokenId), zap.Error(err))
		return
	}
	if used {
		return errors.New("repeated tweet")
	}
	// 获取推文内容
	tweet, err := utils.GetTweetById(s.c, tweetId)
	if err != nil {
		return errors.New("cannot get tweet")
	}
	// 验证推文内容
	if !utils.CheckIfMatchClaimTweet(s.c, req.TokenId, tweet) {
		return errors.New("tweet cannot match")
	}
	// 保存到数据库
	err = s.dao.CreateClaimBadgeTweet(&model.ClaimBadgeTweet{
		Address:    address,
		TokenId:    req.TokenId,
		Url:        req.TweetUrl,
		TweetId:    tweetId,
		AddTs:      time.Now().Unix(),
		Airdropped: false,
	})
	if err != nil {
		log.Errorv("CreateClaimBadgeTweet error", zap.Error(err))
	}
	return
}
