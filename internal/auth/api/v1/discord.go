package v1

import (
	"github.com/gin-gonic/gin"
)

// DiscordAuthorizationURL 获取 Discord 授权链接
func DiscordAuthorizationURL(c *gin.Context) {
	callback := c.Query("callback")
	if data := srv.DiscordAuthorizationURL(callback); data == "" {
		Fail(c)
	} else {
		OkWithData(data, c)
	}
}

// DiscordCallback Discord 回调登陆
func DiscordCallback(c *gin.Context) {
	type DiscordCallback struct {
		Code     string `json:"code" form:"code"`
		Callback string `json:"callback" form:"callback"`
	}
	var discordCallback DiscordCallback
	_ = c.ShouldBindJSON(&discordCallback)

	if id, username, err := srv.DiscordCallback(discordCallback.Code, discordCallback.Callback); err != nil {
		FailWithMessage(err.Error(), c)
	} else {
		OkWithData(map[string]interface{}{"id": id, "username": username}, c)
	}
}
