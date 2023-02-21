package v1

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HashSubmit(c *gin.Context) {
	var submit request.HashSubmitRequest
	_ = c.ShouldBindJSON(&submit)
	address := c.GetString("address")
	if err := service.HashSubmit(address, submit.Hash); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("error", c)
	} else {
		response.Ok(c)
	}
}
