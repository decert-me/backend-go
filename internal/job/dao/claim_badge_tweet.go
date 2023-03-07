package dao

import (
	"backend-go/internal/app/model"
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

func (d *Dao) GetPendingAirdrop() (tokenId []*big.Int, listAddr []string, err error) {
	var pending []model.ClaimBadgeTweet
	if err = d.db.Where("airdropped", false).Find(&pending).Error; err != nil {
		return
	}
	for _, v := range pending {
		tokenId = append(tokenId, big.NewInt(v.TokenId))
		listAddr = append(listAddr, v.Address)
	}
	if len(tokenId) != len(listAddr) {
		err = errors.New("token and address len error")
		return
	}
	return tokenId, listAddr, nil
}

func (d *Dao) UpdateAirdropped(req *model.ClaimBadgeTweet) (err error) {
	err = d.db.Model(&model.ClaimBadgeTweet{}).
		Where("token_id = ? AND address = ?", req.TokenId, req.Address).
		Update("airdropped", true).Error
	return
}

func (d *Dao) UpdateAirdroppedList(tokenIds []*big.Int, receivers []common.Address, hash string) (err error) {
	tx := d.db.Model(&model.ClaimBadgeTweet{}).Begin()
	for i, _ := range receivers {
		tx.Where("token_id = ? AND address = ?", tokenIds[i], receivers[i].String()).
			Updates(map[string]interface{}{"airdropped": "true", "airdrop_hash": hash, "airdrop_ts": time.Now().Unix()})
	}
	return tx.Commit().Error
}
