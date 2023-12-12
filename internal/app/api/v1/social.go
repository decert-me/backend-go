package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetWechatQrcode 获取关注二维码
func GetWechatQrcode(c *gin.Context) {
	c.GetQuery("address")
	address := c.GetString("address")
	if address == "" {
		Fail(c)
		return
	}
	if data, err := srv.GetWechatQrcode(address); err != nil {
		FailWithMessage(err.Error(), c)
	} else {
		OkWithData(data, c)
	}
	return
}

// WechatBindAddress 微信绑定
func WechatBindAddress(c *gin.Context) {
	type WechatBind struct {
		Address string `json:"address" form:"address" binding:"required"`
		Code    string `json:"code" form:"code" binding:"required"`
	}
	var wechatBind WechatBind
	err := c.ShouldBindJSON(&wechatBind)
	if err != nil {
		FailWithMessage("参数错误", c)
		return
	}
	if err := srv.WechatBindAddress(c, wechatBind.Address, wechatBind.Code); err != nil {
		FailWithMessage(err.Error(), c)
	} else {
		Ok(c)
	}
}

// DiscordAuthorizationURL 获取 Discord 授权链接
func DiscordAuthorizationURL(c *gin.Context) {
	callback := c.Query("callback")
	fmt.Println(callback)
	if data, err := srv.DiscordAuthorizationURL(callback); err != nil {
		Fail(c)
	} else {
		OkWithData(data, c)
	}
}

// DiscordBindAddress Discord 回调绑定
func DiscordBindAddress(c *gin.Context) {
	type DiscordCallback struct {
		Code     string `json:"code" form:"code"`
		Callback string `json:"callback" form:"callback"`
	}
	var discordCallback DiscordCallback
	_ = c.ShouldBindJSON(&discordCallback)
	address := c.GetString("address")
	if err := srv.DiscordCallback(address, discordCallback); err != nil {
		FailWithMessage(err.Error(), c)
	} else {
		Ok(c)
	}
}
