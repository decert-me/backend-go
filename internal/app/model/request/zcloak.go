package request

import "gorm.io/datatypes"

type SaveSignAndDidRequest struct {
	Sign       string         `json:"sign" binding:"required"`
	SignHash   string         `gorm:"column:sign_hash;type:varchar(255);comment:签名哈希" json:"sign_hash" form:"sign_hash"`
	DidAddress string         `json:"did_address" binding:"required"`
	KeyFile    datatypes.JSON `json:"key_file" binding:"required"`
}

type GetDidSignMessageRequest struct {
	Did string `form:"did" binding:"required"`
}

type GenerateCardInfoRequest struct {
	TokenId string `json:"token_id" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
	Lang    string `json:"lang"`
}

type GenerateCardRequest struct {
	TokenId string `json:"token_id" binding:"required"`
}

type GetDidCardInfoRequest struct {
	TokenId string `form:"token_id" binding:"required"`
	Address string `form:"address" binding:"required"`
}
