package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

// SaveSignAndDid 保存签名和DID账号
func SaveSignAndDid(c *gin.Context) {
	var request request.SaveSignAndDidRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if err = srv.SaveSignAndDid(address, request); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

// GetAddressDid 查询地址绑定的Did
func GetAddressDid(c *gin.Context) {
	address := c.GetString("address")
	if data, err := srv.GetAddressDid(address); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(map[string]string{"did": data}, c)
	}
}

// GetDidSignMessage 获取 DID 签名
func GetDidSignMessage(c *gin.Context) {
	var request request.GetDidSignMessageRequest
	_ = c.ShouldBindQuery(&request)
	ethAddress := c.GetString("address")
	if loginMessage, err := srv.GetDidSignMessage(request.Did, ethAddress); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(map[string]string{"loginMessage": loginMessage}, GetMessage(c, "FetchSuccess"), c)
	}
}

// GetKeyFileWithSignature 获取KeyFiles签名内容
func GetKeyFileWithSignature(c *gin.Context) {
	address := c.GetString("address")
	if keyFile, err := srv.GetKeyFileWithSignature(address); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(map[string]interface{}{"key_file": keyFile}, c)
	}
}

// GenerateCard 生成证书
func GenerateCard(c *gin.Context) {
	var req request.GenerateCardRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	lang := c.GetString("lang")
	if err = srv.GenerateCard(address, req.TokenId, lang); err != nil {
		Ok(c)
	} else {
		Ok(c)
	}
}

// GetDidCardInfo 获取 DID Card 信息
func GetDidCardInfo(c *gin.Context) {
	var req request.GetDidCardInfoRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if data, err := srv.GetDidCardInfo(req); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}

}
