package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/job/utils"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	twitterClient "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
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
	if err = d.db.Where("status", 0).Where("add_ts < ?", time.Now().Add(-60*time.Second).Unix()).Find(&pending).Error; err != nil {
		return
	}
	for _, v := range pending {
		if v.TweetId == "" {
			// 查找用户推特信息
			var twitterData string
			err = d.db.Raw("SELECT socials->'twitter' FROM users WHERE address = ? LIMIT 1", v.Address).Scan(&twitterData).Error
			if err != nil || twitterData == "" {
				continue
			}
			// 查询用户推特列表
			accessToken := gjson.Get(twitterData, "accessToken").String()
			accessSecret := gjson.Get(twitterData, "accessSecret").String()
			token := oauth1.NewToken(accessToken, accessSecret)
			oaConfig := oauth1.NewConfig(d.c.Auth.Twitter.ConsumerKey, d.c.Auth.Twitter.ConsumerSecret)
			httpClient := oaConfig.Client(oauth1.NoContext, token)
			client := twitterClient.NewClient(httpClient)
			userID := gjson.Get(twitterData, "id").Int()
			list, _, err := client.Timelines.UserTimeline(&twitterClient.UserTimelineParams{UserID: userID, Count: 50})
			if err != nil {
				log.Errorv("List error", zap.Error(err))
				continue
			}
			// 匹配推文
			var match bool
			for _, v2 := range list {
				if utils.CheckIfMatchClaimTweet(d.c, v.TokenId, v2.Text) {
					tokenId = append(tokenId, big.NewInt(v.TokenId))
					listAddr = append(listAddr, v.Address)
					scores = append(scores, big.NewInt(v.Score))
					match = true
					break
				}
			}
			if !match {
				d.UpdateAirdroppedError(v.TokenId, v.Address, "TweetNoMatch")
			}
		} else {
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
