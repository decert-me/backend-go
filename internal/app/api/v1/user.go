package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
	"github.com/gin-gonic/gin"
)

func GetDiscordInfo(c *gin.Context) {
	if list, err := srv.GetDiscordInfo(c.GetString("address")); err != nil {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(list, c)
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
		response.FailWithMessage("address error", c)
		return
	}
	if loginMessage, err := srv.GetLoginMessage(request.Address); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(map[string]string{"loginMessage": loginMessage}, "获取成功", c)
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
		response.FailWithMessage("param not valid", c)
		return
	}
	if token, err := srv.AuthLoginSignRequest(request); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(map[string]string{"token": token}, "获取成功", c)
	}
}
