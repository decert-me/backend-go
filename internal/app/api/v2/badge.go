package v2

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

func SubmitClaimShareV2(c *gin.Context) {
	var req request.SubmitClaimShareV2Req
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	lang := c.GetString("lang")
	if res, err := srv.SubmitClaimShareV2(address, req, lang); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(res, c)
	}
}

// GenerateMintSignature - 生成用户自主 mint NFT 的签名
func GenerateMintSignature(c *gin.Context) {
	var req request.GenerateMintSignatureReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	lang := c.GetString("lang")
	if res, err := srv.GenerateMintSignature(address, req, lang); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(res, c)
	}
}

// ConfirmUserMint - 确认用户自主 mint 成功
func ConfirmUserMint(c *gin.Context) {
	var req request.ConfirmUserMintReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if err := srv.ConfirmUserMint(address, req.TokenId, req.TxHash); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithMessage(GetMessage(c, "Success"), c)
	}
}
