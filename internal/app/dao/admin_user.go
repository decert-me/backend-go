package dao

import (
	"backend-go/internal/app/model"
)

func (d *Dao) IsAdmin(address string) (is bool, err error) {
	var count int64
	if err = d.db.Model(&model.AdminUser{}).Where("address = ?", address).Count(&count).Error; err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
