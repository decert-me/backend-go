package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type QuestTranslated struct {
	gorm.Model
	TokenId     int64          `gorm:"column:token_id;not null;index:language_token_id,unique" json:"tokenId"`
	Title       string         `gorm:"column:title;comment:标题;type:varchar" json:"title" form:"title"` // 标题
	Answer      string         `gorm:"column:answer" json:"answer"`                                    // 答案
	Description string         `gorm:"column:description;type:varchar" json:"description" form:"description"`
	MetaData    datatypes.JSON `gorm:"column:meta_data" json:"metadata"`    // 元数据
	QuestData   datatypes.JSON `gorm:"column:quest_data" json:"quest_data"` // 元数据
	Language    string         `gorm:"column:language;type:varchar(64);not null;index:language_token_id,unique" json:"language"`
}

func (QuestTranslated) TableName() string {
	return "quest_translated"
}
