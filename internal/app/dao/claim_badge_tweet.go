package dao

import (
	"backend-go/internal/app/model"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func (d *Dao) HasTweet(tweetId string) (bool, error) {
	var total int64
	err := d.db.Model(&model.ClaimBadgeTweet{}).
		Where("tweet_id", tweetId).Where("status != 2").
		Count(&total).Error
	return total != 0, err
}

func (d *Dao) CreateClaimBadgeTweet(req *model.ClaimBadgeTweet) (exists bool, err error) {
	var claimd model.ClaimBadgeTweet
	result := d.db.Where("address = ? AND token_id = ?", req.Address, req.TokenId).Where("status = 1").First(&claimd)
	if result.Error == nil {
		return true, nil
	}
	if result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	err = d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_id"}},
		UpdateAll: true,
	}).Create(&req).Error
	return exists, err
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

func (d *Dao) UpdateAirdroppedOne(tokenId int64, receivers string, hash string) (err error) {
	tx := d.db.Model(&model.ClaimBadgeTweet{}).Begin()
	tx.Where("token_id = ? AND address = ?", tokenId, receivers).
		Updates(map[string]interface{}{"status": 1, "airdrop_hash": hash, "airdrop_ts": time.Now().Unix()})
	return tx.Commit().Error
}

func (d *Dao) HasAirdrop(address string, tokenId int64) bool {
	//var total int64
	//err := d.db.Model(&model.ClaimBadgeTweet{}).
	//	Where("address = ? AND token_id = ? AND status=1", address, tokenId).
	//	Count(&total).Error
	//if err != nil {
	//	log.Errorv("HasAirdrop error", zap.Error(err))
	//	return false
	//}
	var total int64
	err := d.db.Model(&model.UserChallenges{}).
		Where("address = ? AND token_id = ?", address, tokenId).
		Count(&total).Error
	if err != nil {
		return false
	}
	return total != 0
}
