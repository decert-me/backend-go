package v1

import (
	"backend-go/internal/auth/model/request"
	"github.com/gin-gonic/gin"
)

// TwitterAuthorizationURL 获取授权链接
func TwitterAuthorizationURL(c *gin.Context) {
	if list, err := srv.TwitterAuthorizationURL(); err != nil {
		Fail(c)
	} else {
		OkWithData(list, c)
	}
}

// TwitterCallback 推特回调登陆
func TwitterCallback(c *gin.Context) {
	var req request.TwitterCallbackReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if result, err := srv.TwitterCallback(req); err != nil {
		Fail(c)
	} else {
		OkWithData(result, c)
	}
}
