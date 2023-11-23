package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

// GetUnreadMessage 获取未读消息
func GetUnreadMessage(c *gin.Context) {
	if list, err := srv.GetUnreadMessage(c.GetString("address")); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(list, c)
	}
}

// ReadMessage 阅读消息
func ReadMessage(c *gin.Context) {
	var r request.ReadMessageRequset
	_ = c.ShouldBindJSON(&r)
	if err := srv.ReadMessage(c.GetString("address"), r.ID); err != nil {
		FailWithMessage(GetMessage(c, "ReadFailed"), c)
	} else {
		OkWithMessage(GetMessage(c, "ReadSuccess"), c)
	}
}
