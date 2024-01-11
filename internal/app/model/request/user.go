package request

import "backend-go/internal/app/model"

type GetLoginMessageRequest struct {
	Address string `json:"address" form:"address"`
}

type AuthLoginSignRequest struct {
	Address   string `json:"address" form:"address" binding:"required"`
	Message   string `json:"message" form:"message" binding:"required"`
	Signature string `json:"signature" form:"signature" binding:"required"`
}

type UpdateUserInfo struct {
	NickName    *string `json:"nickname" form:"nickname"`
	Avatar      *string `json:"avatar" form:"avatar"`
	Description *string `json:"description" form:"description"`
}

type GetUserQuestListRequest struct {
	PageInfo
	model.Quest
	Language string `json:"-" form:"-"`
	Address  string `json:"-" form:"-"`
}
