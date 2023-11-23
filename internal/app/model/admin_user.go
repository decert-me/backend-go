package model

import (
	"gorm.io/gorm"
)

type AdminUser struct {
	gorm.Model
	Username    string `json:"username" gorm:"comment:用户登录名"`
	Address     string `json:"address" gorm:"comment:钱包地址;unique"`
	HeaderImg   string `json:"headerImg" gorm:"comment:用户头像"` // 用户头像
	AuthorityId string `json:"-" gorm:"comment:用户角色ID"`       // 用户角色ID
}

func (AdminUser) TableName() string {
	return "admin_user"
}
