package model

import "time"

type Tag struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `gorm:"column:name;unique;type:varchar(100);comment:标签名" json:"name"`
}

func (Tag) TableName() string {
	return "tag"
}
