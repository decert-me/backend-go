package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Chinese string
	English string
	Weight  int `gorm:"column:weight;default:0"`
}

func (Category) TableName() string {
	return "admin_category"
}
