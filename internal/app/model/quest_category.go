package model

import "gorm.io/gorm"

type QuestCategory struct {
	gorm.Model
	Chinese string
	English string
	Weight  int `gorm:"column:weight;default:0"`
}

func (QuestCategory) TableName() string {
	return "quest_category"
}
