package response

import "backend-go/internal/app/model"

type GetQuestListRes struct {
	model.Quest
	Claimed bool `gorm:"claimed" json:"claimed"`
}
