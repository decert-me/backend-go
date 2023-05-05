package dao

import (
	"backend-go/internal/app/model"
	"gorm.io/gorm/clause"
)

// SetEns 保存Ens信息
func (d *Dao) SetEns(ens model.Ens) (err error) {
	return d.db.Model(&model.Ens{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "domain"}},
		UpdateAll: true,
	}).Create(&ens).Error
}

// GetEnsByAddress 获取Ens信息
func (d *Dao) GetEnsByAddress(address string) (value model.Ens, err error) {
	err = d.db.Model(&model.Ens{}).Where("address", address).First(&value).Error
	return
}

// GetEnsByAddress 获取Ens信息
func (d *Dao) GetAddressByEns(domain string) (value model.Ens, err error) {
	err = d.db.Model(&model.Ens{}).Where("domain", domain).First(&value).Error
	return
}
