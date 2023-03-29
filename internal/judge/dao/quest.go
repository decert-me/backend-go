package dao

import (
	"backend-go/internal/app/model"
)

func (d *Dao) GetQuest(req *model.Quest) (quest model.Quest, err error) {
	err = d.db.Model(&model.Quest{}).Where("token_id", req.TokenId).First(&quest).Error
	return
}
