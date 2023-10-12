package model

import "time"

type Upload struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `json:"-"`                                              // 创建时间
	Address   string    `json:"address" gorm:"column:address;type:varchar(44)"` // 上传人
	Name      string    `json:"name" gorm:"size:50;comment:文件名"`                // 文件名
	Key       string    `json:"key" gorm:"type:char(36);comment:编号"`            // 编号
}
