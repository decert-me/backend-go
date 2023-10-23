package model

import "gorm.io/gorm"

type Theme struct {
	gorm.Model
	Chinese string
	English string
	Weight  int `gorm:"column:weight;default:0;comment:权重"`
}

func (Theme) TableName() string {
	return "admin_theme"
}
