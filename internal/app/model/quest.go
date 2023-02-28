package model

import "gorm.io/datatypes"

type Quest struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	Title        string         `gorm:"column:title;comment:标题;type:varchar" json:"title" form:"title"` // 标题
	Label        string         `gorm:"column:label;comment:标签;type:varchar" json:"-"`                  // 标签
	Disabled     bool           `gorm:"column:disabled" json:"-"`
	Description  string         `gorm:"column:description;type:varchar" json:"description" form:"description"`
	Dependencies datatypes.JSON `gorm:"column:dependencies" json:"-"`
	IsDraft      bool           `gorm:"column:is_draft" json:"-"`
	AddTs        int64          `gorm:"column:add_ts;autoCreateTime:milli" json:"addTs"`
	TokenId      int64          `gorm:"column:token_id;UNIQUE;not null;" json:"tokenId"`
	Type         uint8          `gorm:"column:type" json:"type" form:"type"`        // 0:问答;1:编程
	Difficulty   uint8          `gorm:"column:difficulty" json:"-"`                 // 0:easy;1:moderate;2:difficult
	EstimateTime uint8          `gorm:"column:estimate_time" json:"-"`              // 预估时间/min
	Creator      string         `gorm:"column:creator;type:varchar" json:"creator"` // 用户 address
	MetaData     datatypes.JSON `gorm:"column:meta_data" json:"metadata"`           // 元数据
	ExtraData    datatypes.JSON `gorm:"column:extra_data" json:"-"`                 // 额外数据
	Uri          string         `gorm:"column:uri" json:"uri"`
}
