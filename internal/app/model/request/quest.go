package request

import "backend-go/internal/app/model"

type GetQuestListRequest struct {
	model.Quest
	PageInfo
}

type AddQuestRequest struct {
	Signature   string `json:"signature"`
	Uri         string `json:"uri"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
