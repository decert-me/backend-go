package model

import "gorm.io/datatypes"

type Transaction struct {
	ID        uint           `gorm:"primarykey"`
	AddTs     int64          `gorm:"column:add_ts;autoCreateTime" json:"addTs"`
	Hash      string         `gorm:"column:hash;type:char(68);unique;not null" json:"hash" form:"hash"` // 交易hash唯一
	EventName string         `gorm:"column:event_name;size:20" json:"event_name"`
	SendAddr  string         `gorm:"column:send_addr;type:char(42)" json:"send_addr" form:"send_addr"`
	Raw       string         `gorm:"column:raw" json:"raw" form:"raw"`
	Msg       string         `gorm:"column:msg" json:"msg"`
	Status    uint8          `gorm:"column:status;default:0" json:"status" form:"status"` // 状态 0 处理中 1 交易成功 2 交易失败 3 超过解析次数 4 事件匹配失败 5 出现错误
	Params    datatypes.JSON `gorm:"column:params" json:"params"`
}
