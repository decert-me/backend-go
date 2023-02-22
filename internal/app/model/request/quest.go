package request

import "backend-go/internal/app/model"

type GetQuestListRequest struct {
	model.Quest
	PageInfo
}

type AddQuestRequest struct {
	Signature   string `json:"signature" binding:"required"`
	Uri         string `json:"uri" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
