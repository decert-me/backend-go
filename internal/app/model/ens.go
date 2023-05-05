package model

import "gorm.io/gorm"

type Ens struct {
	gorm.Model
	Address string `gorm:"column:address;index:address_domain,UNIQUE;" json:"address"`
	Domain  string `gorm:"column:domain;index:address_domain,UNIQUE;"  json:"domain"`
	Avatar  string `gorm:"avatar" json:"avatar"`
}
