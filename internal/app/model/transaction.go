package model

import "gorm.io/datatypes"

type Transaction struct {
	ID        uint           `gorm:"primarykey"`
	AddTs     int64          `gorm:"column:add_ts;autoCreateTime;comment:添加时间" json:"addTs"`
	Hash      string         `gorm:"column:hash;type:char(68);unique;not null;comment:交易Hash" json:"hash" form:"hash"` // 交易hash唯一
	EventName string         `gorm:"column:event_name;size:20;comment:事件名称" json:"event_name"`
	SendAddr  string         `gorm:"column:send_addr;type:varchar(44);comment:发送地址" json:"send_addr" form:"send_addr"`
	Raw       string         `gorm:"column:raw" json:"raw" form:"raw"`
	Msg       string         `gorm:"column:msg" json:"msg"`
	Status    uint8          `gorm:"column:status;default:0;comment:交易状态0 处理中 1 交易成功 2 交易失败 3 超过解析次数 4 事件匹配失败 5 出现错误" json:"status" form:"status"` // 状态 0 处理中 1 交易成功 2 交易失败 3 超过解析次数 4 事件匹配失败 5 出现错误
	Params    datatypes.JSON `gorm:"column:params;comment:其它参数" json:"params"`
}
