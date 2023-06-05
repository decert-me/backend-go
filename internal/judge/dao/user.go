package dao

import (
	"backend-go/internal/app/model"
	"time"
)

// GetUserResourceBeforeList 获取活跃时间在之前的地址
func (d *Dao) GetUserResourceBeforeList(before time.Time) (userList []model.Users, err error) {
	err = d.db.Model(&model.Users{}).Where("resource_time < ? AND resource_time > '2023-01-01'", before).Find(&userList).Error
	return
}

func (d *Dao) UpdateUserResourceTime(address string) (err error) {
	err = d.db.Model(&model.Users{}).Where("address = ?", address).Update("resource_time", time.Now()).Error
	return
}
