package dao

import "backend-go/internal/app/model"

func (d *Dao) SaveJudgeResult(res model.JudgeResult) (id string, err error) {
	if err = d.db.Create(&res).Error; err != nil {
		return
	}
	return res.ID, nil
}
