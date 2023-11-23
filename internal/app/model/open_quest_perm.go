package model

import "time"

type OpenQuestPerm struct {
	ID        int64 `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time
	Address   string `gorm:"column:address;type:varchar(44);comment:钱包地址;UNIQUE" json:"address" form:"address"`
}
