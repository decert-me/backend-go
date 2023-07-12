package v1

import (
	"backend-go/internal/app/model/request"
	v1 "backend-go/internal/judge/api/v1"
	"github.com/gin-gonic/gin"
)

func GetProgress(c *gin.Context) {
	var req request.GetProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	userID := c.GetUint("userID")
	if userID == 0 {
		FailWithDetailed(gin.H{"reload": true}, v1.GetMessage(c, "UnauthorizedAccess"), c)
		return
	}
	if data, err := srv.GetProgress(userID, req); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

func UpdateProgress(c *gin.Context) {
	var req request.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	userID := c.GetUint("userID")
	if userID == 0 {
		FailWithDetailed(gin.H{"reload": true}, v1.GetMessage(c, "UnauthorizedAccess"), c)
		return
	}
	if err := srv.UpdateProgress(userID, req); err != nil {
		FailWithMessage(GetMessage(c, "UpdateFailed"), c)
	} else {
		Ok(c)
	}
}
