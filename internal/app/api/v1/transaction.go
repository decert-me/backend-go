package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

func HashSubmit(c *gin.Context) {
	var submit request.HashSubmitRequest
	_ = c.ShouldBindJSON(&submit)
	address := c.GetString("address")
	if err := srv.HashSubmit(address, submit); err != nil {
		FailWithMessage(GetMessage(c, "OperationFailed"), c)
	} else {
		Ok(c)
	}
}
