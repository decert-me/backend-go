package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/job/utils"
	"fmt"

	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

func (d *Dao) HasTweet(tweetId string) (bool, error) {
	var total int64
	err := d.db.Model(&model.ClaimBadgeTweet{}).
		Where("tweet_id", tweetId).
		Count(&total).Error
	return total != 0, err
}

func (d *Dao) CreateClaimBadgeTweet(req *model.ClaimBadgeTweet) (err error) {
	return d.db.Create(req).Error
}

func (d *Dao) GetPendingAirdrop() (tokenId []*big.Int, listAddr []string, scores []*big.Int, err error) {
	var pending []model.ClaimBadgeTweet
	if err = d.db.Where("status", 0).Find(&pending).Error; err != nil {
		return
	}
	for _, v := range pending {
		// 获取推文内容
		tweet, err := utils.GetTweetById(d.c, v.TweetId)
		if err != nil {
			d.UpdateAirdroppedError(v.TokenId, v.Address, fmt.Sprintf("GetTweetById err:%s", err.Error()))
			continue
		}
		// 验证推文内容
		if !utils.CheckIfMatchClaimTweet(d.c, v.TokenId, tweet) {
			d.UpdateAirdroppedError(v.TokenId, v.Address, "InconsistentTweet")
			continue
		}
		tokenId = append(tokenId, big.NewInt(v.TokenId))
		listAddr = append(listAddr, v.Address)
		scores = append(scores, big.NewInt(v.Score))
	}
	if len(tokenId) != len(listAddr) {
		err = errors.New("token and address len error")
		return
	}
	return tokenId, listAddr, scores, nil
}

func (d *Dao) UpdateAirdropped(req *model.ClaimBadgeTweet) (err error) {
	err = d.db.Model(&model.ClaimBadgeTweet{}).
		Where("token_id = ? AND address = ?", req.TokenId, req.Address).
		Update("status", 1).Error
	return
}

func (d *Dao) UpdateAirdroppedList(tokenIds []*big.Int, receivers []common.Address, hash string) (err error) {
	tx := d.db.Model(&model.ClaimBadgeTweet{}).Begin()
	for i, _ := range receivers {
		tx.Where("token_id = ? AND address = ?", tokenIds[i], receivers[i].String()).
			Updates(map[string]interface{}{"status": 1, "airdrop_hash": hash, "airdrop_ts": time.Now().Unix()})
	}
	return tx.Commit().Error
}

func (d *Dao) UpdateAirdroppedError(tokenId int64, address string, msg string) (err error) {
	err = d.db.Model(&model.ClaimBadgeTweet{}).
		Where("token_id = ? AND address = ?", tokenId, address).
		Updates(map[string]interface{}{"msg": msg, "status": 2}).Error
	return
}
