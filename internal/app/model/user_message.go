package model

import "gorm.io/gorm"

type UserMessage struct {
	gorm.Model
	Address   string `gorm:"column:address;type:varchar(44);comment:钱包地址;index:address_is_read" json:"address" form:"address"`
	TokenId   int64  `gorm:"column:token_id" json:"token_id"`
	Title     string `gorm:"column:title;type:varchar(200);default:''" json:"title" form:"title"`                           // 消息标题
	TitleEn   string `gorm:"column:title_en;type:varchar(200);default:''" json:"title_en" form:"title_en"`                  // 消息标题
	Content   string `gorm:"column:content;type:text;default:'';comment:消息内容" json:"content" form:"content"`                // 消息内容
	ContentEn string `gorm:"column:content_en;type:text;default:'';comment:消息内容" json:"content_en" form:"content_en"`       // 消息内容
	IsRead    bool   `gorm:"column:is_read;default:false;comment:是否已读;index:address_is_read" json:"is_read" form:"is_read"` // 是否已读
}
