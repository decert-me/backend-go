package service

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"go.uber.org/zap"
	"math/big"
	"time"
)

func (s *Service) PermitClaimBadge(address string, req request.PermitClaimBadgeReq) (res response.PermitClaimBadgeRes, err error) {
	// 校验分数正确性
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenId})
	if err != nil {
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
	chain := s.IDToMultiChain[req.ChainID]
	if chain.Badge == "" {
		return res, errors.New("UnSupportChain")
	}
	// TODO: 判断状态
	status, err := s.dao.GetMultiChainStatus(req.TokenId, req.ChainID)
	if err != nil {
		return res, errors.New("UnSupportChain")
	}
	var hash []byte
	if status == 0 {
		// TODO: 从链上获取quest是否更改
		var questData model.Extradata
		err = json.Unmarshal(quest.ExtraData, &questData)
		if err != nil {
			log.Errorv("Unmarshal error", zap.Error(err))
			return res, errors.New("UnexpectedError")
		}
		hash = solsha3.SoliditySHA3(
			[]string{"address", "uint256", "uint32", "uint32", "string", "string", "address", "string", "address", "address"},
			[]interface{}{
				questData.Creator, big.NewInt(req.TokenId), questData.StartTs, questData.EndTs, questData.Title, questData.Uri, req.To, "uri", badgeAddress, address,
			},
		)
		res.Func = "claimWithInit"
	} else {
		hash = solsha3.SoliditySHA3(
			[]string{"address", "uint256", "string", "address", "address"},
			[]interface{}{
				req.To, big.NewInt(req.TokenId), "uri", chain.Badge, address,
			},
		)
		res.Func = "claim"
	}
	res.Uri = "uri"
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	res.Sign = hexutil.Encode(signature)
	return res, err
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
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenId})
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

func (s *Service) ShouldInitQuest(chain config.MultiChain) {

}
