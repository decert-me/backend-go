package dao

import "backend-go/internal/app/model"

func (d *Dao) FilterJudgeResult(judgeResult model.JudgeResult) (Judge model.JudgeResult, err error) {
	err = d.db.Model(&model.JudgeResult{}).Where(judgeResult).First(&Judge).Error
	return
}
