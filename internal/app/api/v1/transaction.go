package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

func HashSubmit(c *gin.Context) {
	var submit request.HashSubmitRequest
	_ = c.ShouldBindJSON(&submit)
	address := c.GetString("address")
	if err := srv.HashSubmit(address, submit.Hash); err != nil {
		response.FailWithMessage("操作失败", c)
	} else {
		response.Ok(c)
	}
}
