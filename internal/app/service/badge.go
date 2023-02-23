package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (s *Service) PermitClaimBadge(address string, req request.PermitClaimBadgeReq) (res string, err error) {
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
	if err != nil {
		s.log.Error("ValidTokenId error", zap.Error(err))
	}
	if !valid {
		return errors.New("invalid quest")
	}
	// TODO: 检查用户是否已通过挑战，或已领取SBT

	// 获取推文ID
	tweetId := utils.GetTweetIdFromURL(req.TweetUrl)
	if tweetId == "" {
		return errors.New("cannot find tweet id'")
	}
	// 检查是否重复使用
	used, err := s.dao.HasTweet(tweetId)
	if err != nil {
		s.log.Error("HasTweet error", zap.Error(err))
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
	return
}
