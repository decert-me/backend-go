package model

import "github.com/lib/pq"

type OpenQuestUserPerm struct {
	ID          int64          `gorm:"column:id;primary_key" json:"id"`
	QuestID     uint           `gorm:"column:quest_id" json:"quest_id"`
	AddressList pq.StringArray `gorm:"column:address_list;type:varchar[]" json:"address_list"`
}
