package model

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Chinese string
	English string
	Weight  int `gorm:"column:weight;default:0"`
}

func (Language) TableName() string {
	return "admin_language"
}
