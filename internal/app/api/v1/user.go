package v1

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/service"
	"backend-go/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDiscordInfo(c *gin.Context) {
	if list, err := service.GetDiscordInfo(c.GetString("address")); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
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
	if err, loginMessage := service.GetLoginMessage(request.Address); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
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
	_ = c.ShouldBindJSON(&request)
	if token, err := service.AuthLoginSignRequest(request); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(map[string]string{"token": token}, "获取成功", c)
	}
}
