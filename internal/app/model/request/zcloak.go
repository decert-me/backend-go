package request

import "gorm.io/datatypes"

type SaveSignAndDidRequest struct {
	Sign       string         `json:"sign" binding:"required"`
	DidAddress string         `json:"did_address" binding:"required"`
	KeyFile    datatypes.JSON `json:"key_file" binding:"required"`
}

type GetDidSignMessageRequest struct {
	Did string `form:"did" binding:"required"`
}

type GenerateCardInfoRequest struct {
	TokenId int64  `json:"token_id"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
}
