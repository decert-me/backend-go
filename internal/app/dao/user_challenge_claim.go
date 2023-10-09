package dao

import (
	"backend-go/internal/app/model"
	"gorm.io/gorm"
)

// CreateUserChallengeClaim 创建记录
func (d *Dao) CreateUserChallengeClaim(claim *model.UserChallengeClaim) (err error) {
	err = d.db.Create(&claim).Error
	return
}

// HasClaimed 查询是否claim
func (d *Dao) HasClaimed(address string, tokenID int64) (status uint8, err error) {
	var claim model.UserChallengeClaim
	err = d.db.Model(&model.UserChallengeClaim{}).Where("address = ? AND token_id = ?", address, tokenID).First(&claim).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return claim.Status, nil
}
