package dao

import (
	"backend-go/internal/app/model"
	"gorm.io/gorm"
)

// HasOpenQuestPerm 获取用户是否有开放题权限
func (d *Dao) HasOpenQuestPerm(address string) (perm bool, beta bool, err error) {
	// 获取系统设置
	var betaSet string
	err = d.db.Model(&model.SystemSetting{}).Select("COALESCE(value,'false')").Where("key = 'beta'").Scan(&betaSet).Error
	if err != nil {
		return
	}
	if betaSet == "true" {
		beta = true
	}
	// 查询是否管理员
	if err := d.db.Model(&model.AdminUser{}).Where("address ILIKE ?", address).First(&model.AdminUser{}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return false, beta, err
		}
	} else {
		return true, beta, err
	}

	err = d.db.Model(&model.OpenQuestPerm{}).
		Where("address ILIKE ?", address).
		First(&model.OpenQuestPerm{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, beta, nil
		}
		return false, beta, err
	}
	return true, beta, err
}
