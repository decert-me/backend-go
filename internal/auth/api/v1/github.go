package v1

import "github.com/gin-gonic/gin"

// GithubAuthorizationURL 获取授权链接
func GithubAuthorizationURL(c *gin.Context) {
	callback := c.Query("callback")
	if data, err := srv.GithubAuthorizationURL(callback); err != nil {
		Fail(c)
	} else {
		OkWithData(data, c)
	}
}

// GithubCallback 推特回调登陆
func GithubCallback(c *gin.Context) {
	type GithubCallback struct {
		Code     string `json:"code" form:"code"`
		Callback string `json:"callback" form:"callback"`
	}
	var githubCallback GithubCallback
	_ = c.ShouldBindJSON(&githubCallback)
	code := githubCallback.Code
	callback := githubCallback.Callback
	if id, username, err := srv.GithubCallback(code, callback); err != nil {
		Fail(c)
	} else {
		OkWithData(map[string]interface{}{"id": id, "username": username}, c)
	}
}
