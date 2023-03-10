package request

import "backend-go/internal/app/model"

type GetQuestListRequest struct {
	model.Quest
	PageInfo
	Address string
}

type AddQuestRequest struct {
	Uri         string `json:"uri" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartTs     string `json:"start_ts" binding:"required"`
	EndTs       string `json:"end_ts" binding:"required"`
	Supply      string `json:"supply" binding:"required"`
}
