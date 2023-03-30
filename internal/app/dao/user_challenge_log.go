package dao

import "backend-go/internal/app/model"

func (d *Dao) CreateChallengeLog(log *model.UserChallengeLog) (err error) {
	return d.db.Create(&log).Error
}
