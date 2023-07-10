package v1

import "github.com/gin-gonic/gin"

// GithubAuthorizationURL 获取授权链接
func GithubAuthorizationURL(c *gin.Context) {
	if list, err := srv.GithubAuthorizationURL(); err != nil {
		Fail(c)
	} else {
		OkWithData(list, c)
	}
}

// GithubCallback 推特回调登陆
func GithubCallback(c *gin.Context) {
	//if list, err := srv.TwitterAuthorizationURL(); err != nil {
	//	Fail(c)
	//} else {
	//	OkWithData(list, c)
	//}
}
