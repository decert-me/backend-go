package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ZcloakCard struct {
	gorm.Model
	AddTs   int64          `gorm:"column:add_ts;autoCreateTime" json:"addTs"`
	Address string         `gorm:"column:address;type:varchar(100);comment:钱包地址;index:address_quest_id" json:"address" form:"address"`
	Did     string         `gorm:"column:did" json:"did"`
	QuestID uint           `gorm:"column:quest_id;index:address_quest_id" json:"quest_id"`
	Score   int64          `gorm:"column:score" json:"score"` // 分数
	VC      datatypes.JSON `gorm:"column:vc" json:"vc"`
}
