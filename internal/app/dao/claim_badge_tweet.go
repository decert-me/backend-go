package dao

import (
	"backend-go/internal/app/model"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (d *Dao) HasTweet(tweetId string) (bool, error) {
	var total int64
	err := d.db.Model(&model.ClaimBadgeTweet{}).
		Where("tweetId", tweetId).
		Count(&total).Error
	return total != 0, err
}

func (d *Dao) CreateClaimBadgeTweet(req *model.ClaimBadgeTweet) (err error) {
	return d.db.Create(req).Error
}

func (d *Dao) GetPendingAirdrop() (res map[int64][]string, err error) {
	res = make(map[int64][]string)
	var pending []model.ClaimBadgeTweet
	if err = d.db.Where("airdropped", false).Find(&pending).Error; err != nil {
		return
	}
	for _, v := range pending {
		res[v.TokenId] = append(res[v.TokenId], v.Address)
	}
	return res, nil
}

func (d *Dao) UpdateAirdropped(req *model.ClaimBadgeTweet) (err error) {
	err = d.db.Model(&model.ClaimBadgeTweet{}).
		Where("token_id = ? AND address = ?", req.TokenId, req.Address).
		Update("airdropped", true).Error
	return
}

func (d *Dao) UpdateAirdroppedList(tokenId int64, receivers []common.Address, hash string) (err error) {
	tx := d.db.Model(&model.ClaimBadgeTweet{}).Begin()
	for _, v := range receivers {
		tx.Where("token_id = ? AND address = ?", tokenId, v.String()).
			Updates(map[string]interface{}{"airdropped": "true", "airdrop_hash": hash, "airdrop_ts": time.Now().Unix()})
	}
	return tx.Commit().Error
}
