package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"github.com/gin-gonic/gin"
)

func GetDiscordInfo(c *gin.Context) {
	if list, err := srv.GetDiscordInfo(c.GetString("address")); err != nil {
		FailWithMessage(GetMessage(c, "NoBindingDetected"), c)
	} else {
		OkWithData(list, c)
	}
}

// GetLoginMessage
// @Tags SignApi
// @Summary 获取登录签名消息
// @accept application/json
// @Produce application/json
// @Router /sign/getLoginMessage [post]
func GetLoginMessage(c *gin.Context) {
	var request request.GetLoginMessageRequest
	_ = c.ShouldBindQuery(&request)
	if !utils.IsValidAddress(request.Address) {
		FailWithMessage(GetMessage(c, "AddressError"), c)
		return
	}
	if loginMessage, err := srv.GetLoginMessage(request.Address); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(map[string]string{"loginMessage": loginMessage}, GetMessage(c, "FetchSuccess"), c)
	}
}

// AuthLoginSign
// @Tags SignApi
// @Summary 校验登录签名
// @accept application/json
// @Produce application/json
// @Router /sign/authLoginSign [post]
func AuthLoginSign(c *gin.Context) {
	var request request.AuthLoginSignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if token, err := srv.AuthLoginSignRequest(request); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithDetailed(map[string]string{"token": token}, GetMessage(c, "FetchSuccess"), c)
	}
}
