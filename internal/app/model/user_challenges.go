package model

import (
	"gorm.io/datatypes"
)

type UserChallenges struct {
	ID       uint           `gorm:"primarykey"`
	Address  string         `gorm:"column:address;type:char(42);index:address_questId,UNIQUE;comment:钱包地址" json:"address" form:"address"`
	QuestID  uint           `gorm:"column:questId;index:address_questId,UNIQUE;"`
	Status   uint8          `gorm:"column:status;default:0;size:30;comment:0 进行中 1 等待验证 2 成功;" json:"status" form:"status"` // 0:进行中;1:等待验证;2:成功;
	Content  datatypes.JSON `gorm:"" json:"content"`
	AddTs    int64          `gorm:"autoCreateTime" json:"addTs"`
	Claimed  bool           `gorm:"default:false" json:"claimed"`
	UpdateTs int64          `gorm:"autoUpdateTime" json:"updateTs"`
	ClaimTs  int64          `gorm:"" json:"claimTs"`
}
