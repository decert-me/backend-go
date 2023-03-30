package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

func CreateChallengeLog(c *gin.Context) {
	var add request.SaveChallengeLogRequest
	if err := c.ShouldBindJSON(&add); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	add.Address = c.GetString("address")
	if err := srv.CreateChallengeLog(add); err != nil {
		Fail(c)
	} else {
		Ok(c)
	}
}
