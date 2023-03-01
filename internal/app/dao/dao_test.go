package dao

import (
	"backend-go/internal/app/model"
	"gorm.io/gorm"
)

func (d *Dao) DeleteQuest() {
	err := d.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Quest{}).Error
	if err != nil {
		panic(err)
	}
}

func (d *Dao) DeleteChallenges() {
	err := d.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.UserChallenges{}).Error
	if err != nil {
		panic(err)
	}
}

func (d *Dao) DeleteBadgeTweet() {
	err := d.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.ClaimBadgeTweet{}).Error
	if err != nil {
		panic(err)
	}
}
