package dao

import (
	"backend-go/internal/app/model"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
)

func (d *Dao) TwitterIsBinding(userID int64) (bool, error) {
	var count int
	err := d.db.Raw("SELECT count(1) FROM users WHERE socials->'twitter'->>'id' = ?", strconv.Itoa(int(userID))).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, err
}

func (d *Dao) TwitterBinding(address string, userID int64, userName string) error {
	var count int
	userInfo := fmt.Sprintf("{\"id\": \"%s\", \"username\": \"%s\"}", strconv.Itoa(int(userID)), userName)
	err := d.db.Raw("UPDATE users SET socials = jsonb_set(socials, '{\"twitter\"}', ?, true) WHERE address = ?", userInfo, address).Scan(&count).Error
	return err
}

func (d *Dao) TwitterQueryIdByAddress(address string) (twitterID string, err error) {
	err = d.db.Raw("SELECT socials->'twitter'->>'id' FROM users WHERE address = ? LIMIT 1", address).Scan(&twitterID).Error
	if err != nil {
		return twitterID, err
	}
	return twitterID, err
}

func (d *Dao) TwitterCreateTweetClaim(req *model.ClaimBadgeTweet) (exists bool, err error) {
	var claim model.ClaimBadgeTweet
	result := d.db.Where("address = ? AND token_id = ?", req.Address, req.TokenId).Where("status = 1").First(&claim)
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
