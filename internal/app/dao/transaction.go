package dao

import (
	"backend-go/internal/app/model"
)

func (d *Dao) CreateTransaction(req *model.Transaction) (err error) {
	err = d.db.Model(&model.Transaction{}).Create(req).Error
	return

}
