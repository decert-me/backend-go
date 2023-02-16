package model

import "gorm.io/datatypes"

type Quest struct {
	ID           uint           `gorm:"primarykey"`
	Title        string         `gorm:"column:title;comment:标题;type:varchar" json:"title" form:"title"` // 标题
	Label        string         `gorm:"column:label;comment:标签;type:varchar" json:"label" form:"label"` // 标签
	Disabled     bool           `gorm:"column:disabled" json:"disabled" form:"disabled"`
	Description  string         `gorm:"column:description;type:varchar" json:"description" form:"description"`
	Dependencies datatypes.JSON `gorm:"dependencies" json:"dependencies"`
	IsDraft      bool           `gorm:"column:is_draft" json:"isDraft"`
	AddTs        int64          `gorm:"autoCreateTime" json:"addTs"`
	TokenId      string         `gorm:"column:token_id;type:varchar" json:"token"`
	Type         uint8          `gorm:"column:type" json:"type" form:"type"`                   // 0:问答;1:编程
	Difficulty   uint8          `gorm:"column:difficulty" json:"difficulty" form:"difficulty"` // 0:easy;1:moderate;2:difficult
	EstimateTime uint8          `gorm:"column:estimate_time" json:"estimateTime"`              // 预估时间/min
	Creator      string         `gorm:"column:token_id;type:varchar" json:"token"`             // 用户 address
	MetaData     datatypes.JSON `gorm:"" json:"metadata"`                                      // 元数据
	ExtraData    datatypes.JSON `gorm:"" json:"extradata"`                                     // 额外数据
	Uri          string         `gorm:"column:uri" json:"uri"`
}
