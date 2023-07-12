package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

func GetProgress(c *gin.Context) {
	var req request.GetProgressRequest
	_ = c.ShouldBindQuery(&req)
	userID := c.GetUint("userID")
	if data, err := srv.GetProgress(userID, req); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

func UpdateProgress(c *gin.Context) {
	var req request.UpdateProgressRequest
	_ = c.ShouldBindQuery(&req)
	userID := c.GetUint("userID")
	if err := srv.UpdateProgress(userID, req); err != nil {
		FailWithMessage(GetMessage(c, "UpdateFailed"), c)
	} else {
		Ok(c)
	}
}
