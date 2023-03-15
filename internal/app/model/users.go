package model

import (
	"gorm.io/datatypes"
)

type Users struct {
	ID                uint           `gorm:"primarykey" json:"-"`
	Address           string         `gorm:"column:address;type:char(42);UNIQUE;comment:钱包地址" json:"address" form:"address"`
	NickName          *string        `gorm:"column:nickname;type:varchar(200);default:''" json:"nickname" form:"nickname"`
	Avatar            *string        `gorm:"column:avatar;type:varchar(200);comment:用户头像;default:''" json:"avatar" form:"avatar"`
	Description       *string        `gorm:"column:description;type:varchar(100);comment:自我介绍;default:''" json:"description" form:"description"`
	CreationTimestamp int64          `gorm:"column:creation_timestamp;autoCreateTime" json:"creationTimestamp"`
	Socials           datatypes.JSON `gorm:"column:socials" json:"socials"`
}
