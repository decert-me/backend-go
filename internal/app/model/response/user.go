package response

import "backend-go/internal/app/model"

type GetUserQuestListRes struct {
	model.Quest
}

type QuestWithClaimed struct {
	HasClaim bool `gorm:"column:has_claim" json:"has_claim"`
	model.Quest
}

func (QuestWithClaimed) TableName() string {
	return "quest"
}
