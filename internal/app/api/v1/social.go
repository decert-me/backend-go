package v1

import (
	"backend-go/internal/app/model/request"
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
		FailWithMessage(GetMessage(c, err.Error()), c)
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
		FailWithMessage(GetMessage(c, "ParameterError"), c)
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
	if data, err := srv.DiscordAuthorizationURL(callback); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
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
		OkWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithMessage("", c)
	}
}

// GetEmailBindCode 获取邮箱绑定验证码
func GetEmailBindCode(c *gin.Context) {
	var r request.GetEmailBindCodeRequest
	if c.ShouldBindJSON(&r) != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if address == "" {
		FailWithMessage(GetMessage(c, "SignatureExpired"), c)
		return
	}
	if err := srv.GetEmailBindCode(address, r.Email); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

// EmailBindAddress 绑定邮箱
func EmailBindAddress(c *gin.Context) {
	var r request.EmailBindAddressRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		fmt.Println(err)
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if address == "" {
		FailWithMessage(GetMessage(c, "SignatureExpired"), c)
		return
	}
	if err := srv.EmailBindAddress(address, r.Email, r.Code); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

// GithubAuthorizationURL 获取 Github 授权链接
func GithubAuthorizationURL(c *gin.Context) {
	callback := c.Query("callback")
	if data, err := srv.GithubAuthorizationURL(callback); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(data, c)
	}
}

// GithubBindAddress Github 回调绑定
func GithubBindAddress(c *gin.Context) {
	type GithubCallback struct {
		Code     string `json:"code" form:"code"`
		Callback string `json:"callback" form:"callback"`
	}
	var githubCallback GithubCallback
	_ = c.ShouldBindJSON(&githubCallback)
	address := c.GetString("address")
	if err := srv.GithubCallback(address, githubCallback); err != nil {
		OkWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithMessage("", c)
	}
}
