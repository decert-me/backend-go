package request

import (
	"gorm.io/datatypes"
	"time"
)

type GetUserOpenQuestListRequest struct {
	PageInfo
	OpenQuestReviewStatus uint8 `json:"open_quest_review_status" form:"open_quest_review_status"`
}

type ReviewOpenQuestRequest struct {
	ID        uint           `json:"id" binding:"required"`
	Answer    datatypes.JSON `json:"answer" binding:"required"`
	Index     int            `json:"index" binding:"required"`
	UpdatedAt *time.Time     `json:"updated_at" binding:"required"`
}

type GetUserOpenQuestDetailListRequest struct {
	PageInfo
	TokenID               string `json:"token_id" binding:"required"`
	Index                 *uint8 `json:"index" binding:"required"`
	OpenQuestReviewStatus uint8  `json:"open_quest_review_status" form:"open_quest_review_status"`
}
