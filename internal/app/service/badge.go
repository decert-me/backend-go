package service

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"time"
)

func PermitClaimBadge(address string, req request.PermitClaimBadgeReq) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(global.CONFIG.BlockChain.PrivateKey)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	//hash := solsha3.SoliditySHA3(
	//	[]string{"uint256", "uint256", "address", "address"},
	//	[]interface{}{
	//		big.NewInt(req.TokenId), big.NewInt(req.Score), global.CONFIG.Contract.Badge, address,
	//	},
	//)
	hash := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "address"},
		[]interface{}{
			big.NewInt(req.TokenId), global.CONFIG.Contract.Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func SubmitClaimTweet(address string, req request.SubmitClaimTweetReq) (err error) {
	// 检查tokenId是否存在以及可用
	err = global.DB.Model(&model.Quest{}).
		Where("tokenId", req.TokenId).Where("disabled", false).Where("isDraft", false).
		First(&model.Quest{}).Error
	if err != nil {
		return errors.New("invalid quest")
	}
	// TODO: 检查用户是否已通过挑战，或已领取SBT

	// 获取推文ID
	tweetId := utils.GetTweetIdFromURL(req.TweetUrl)
	if tweetId == "" {
		return errors.New("cannot find tweet id'")
	}
	// 检查是否重复使用
	var count int64
	err = global.DB.Model(&model.ClaimBadgeTweet{}).
		Where("tweetId", tweetId).
		Count(&count).Error
	if count > 0 {
		return errors.New("repeated tweet")
	}
	// 获取推文内容
	tweet, err := utils.GetTweetById(tweetId)
	if err != nil {
		return errors.New("cannot get tweet")
	}
	// 验证推文内容
	if !utils.CheckIfMatchClaimTweet(req.TokenId, tweet) {
		return errors.New("tweet cannot match")
	}
	// 保存到数据库
	claimBadgeTweet := model.ClaimBadgeTweet{
		Address:   address,
		TokenId:   req.TokenId,
		Url:       req.TweetUrl,
		TweetId:   tweetId,
		AddTs:     time.Now().Unix(),
		AirDroped: true,
	}
	err = global.DB.Create(&claimBadgeTweet).Error
	return
}
