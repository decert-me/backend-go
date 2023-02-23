package dao

import (
	"backend-go/internal/app/model"
	"gorm.io/gorm/clause"
)

func (d *Dao) CreateChallenges(challenges *model.UserChallenges) (err error) {
	err = d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "questId"}},
		UpdateAll: true,
	}).Create(&challenges).Error
	return
}
