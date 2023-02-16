package model

import (
	"gorm.io/datatypes"
)

type Users struct {
	ID                uint           `gorm:"primarykey"`
	Address           string         `gorm:"column:address;type:char(42);UNIQUE;comment:钱包地址" json:"address" form:"address"`
	CreationTimestamp int64          `gorm:"autoCreateTime" json:"creationTimestamp"`
	Socials           datatypes.JSON `gorm:"" json:"socials"`
}
