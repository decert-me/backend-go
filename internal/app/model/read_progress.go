package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ReadProgress struct {
	gorm.Model
	UserID        uint           `gorm:"column:user_id" json:"-"`
	CatalogueName string         `gorm:"column:catalogue_name" json:"catalogue_name" form:"catalogue_name"`
	Data          datatypes.JSON `gorm:"column:data" json:"data"` // 数据
	Hash          string         `gorm:"column:hash" json:"-"`    // hash
}
