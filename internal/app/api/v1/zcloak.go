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

// GetVcInfo 获取 VC 信息
func GetVcInfo(c *gin.Context) {
	id := c.GetString("id")
	address := c.Param("address")
	if data, err := srv.GetVcInfo(address, id); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

// GetDidSignMessage 获取 DID 签名
func GetDidSignMessage(c *gin.Context) {
	var request request.GetDidSignMessageRequest
	_ = c.ShouldBindQuery(&request)
	if loginMessage, nonce, err := srv.GetDidSignMessage(request.Did); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(map[string]string{"loginMessage": loginMessage, "nonce": nonce}, GetMessage(c, "FetchSuccess"), c)
	}
}

// GenerateCardInfo 生成卡片信息
func GenerateCardInfo(c *gin.Context) {
	var req request.GenerateCardInfoRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if err = srv.GenerateCardInfo(address, 0, req); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

// GetKeyFileWithSignature 获取KeyFiles签名内容
func GetKeyFileWithSignature(c *gin.Context) {
	address := c.GetString("address")
	if data, keyFile, nonce, err := srv.GetKeyFileWithSignature(address); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(map[string]interface{}{"key_file": keyFile, "signature": data, "nonce": nonce}, c)
	}
}
