package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ZcloakDid struct {
	gorm.Model
	SignMessage string         `gorm:"column:sign_message;type:varchar(255);comment:签名消息" json:"sign_message" form:"sign_message"`
	Signature   string         `gorm:"column:signature;type:varchar(255);comment:签名" json:"signature" form:"signature"`
	Address     string         `gorm:"column:address;type:varchar(100);comment:钱包地址" json:"address" form:"address"`
	DidAddress  string         `gorm:"column:did_address;type:varchar(100);comment:DID地址" json:"did_address" form:"did_address"`
	KeyFile     datatypes.JSON `gorm:"column:key_file" json:"key_file"`
}
