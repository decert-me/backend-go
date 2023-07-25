package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

// ClickShare 点击分享
func ClickShare(c *gin.Context) {
	var submit request.ClickShareRequest
	_ = c.ShouldBindJSON(&submit)
	if err := srv.ClickShare(c, submit); err != nil {
		FailWithMessage(GetMessage(c, "OperationFailed"), c)
	} else {
		Ok(c)
	}
}

// AirdropCallback 空投回调
func AirdropCallback(c *gin.Context) {
	var submit request.AirdropCallbackRequest
	_ = c.ShouldBindJSON(&submit)
	if err := srv.AirdropCallback(submit); err != nil {
		FailWithMessage(GetMessage(c, "OperationFailed"), c)
	} else {
		Ok(c)
	}
}
