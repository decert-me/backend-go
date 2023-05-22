package v1

import (
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
	if list, err := srv.TwitterAuthorizationURL(); err != nil {
		Fail(c)
	} else {
		OkWithData(list, c)
	}
}
