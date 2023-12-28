package model

import (
	"gorm.io/gorm"
)

type CollectionTranslated struct {
	gorm.Model
	CollectionID int64  `gorm:"column:collection_id;not null;comment:合辑ID;index:language_collection_id,UNIQUE" json:"collection_id"`
	Title        string `gorm:"column:title;comment:标题;type:varchar" json:"title" form:"title"` // 标题
	Description  string `gorm:"column:description;type:varchar" json:"description" form:"description"`
	Language     string `gorm:"column:language;type:varchar(64);not null;index:language_collection_id,unique" json:"language"`
}

func (CollectionTranslated) TableName() string {
	return "collection_translated"
}
