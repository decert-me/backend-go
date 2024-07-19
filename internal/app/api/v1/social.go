package v1

import (
	"backend-go/internal/app/model/request"
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
		Replace bool   `json:"replace" form:"replace"`
	}
	var wechatBind WechatBind
	err := c.ShouldBindJSON(&wechatBind)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if err := srv.WechatBindAddress(c, wechatBind.Address, wechatBind.Code, wechatBind.Replace); err != nil {
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
		Replace  bool   `json:"replace" form:"replace"`
	}
	var discordCallback DiscordCallback
	_ = c.ShouldBindJSON(&discordCallback)
	address := c.GetString("address")
	if data, err := srv.DiscordCallback(address, discordCallback, discordCallback.Replace); err != nil {
		OkWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(data, c)
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
	language := c.GetString("lang")
	if err := srv.GetEmailBindCode(address, r.Email, language); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

// EmailBindAddress 绑定邮箱
func EmailBindAddress(c *gin.Context) {
	var r request.EmailBindAddressRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if address == "" {
		FailWithMessage(GetMessage(c, "SignatureExpired"), c)
		return
	}
	if data, err := srv.EmailBindAddress(address, r.Email, r.Code, r.Replace); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		message := "操作成功"
		if data.Success == false {
			message = "绑定失败"
		}
		OkWithDetailed(data, message, c)
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
		Replace  bool   `json:"replace" form:"replace"`
	}
	var githubCallback GithubCallback
	_ = c.ShouldBindJSON(&githubCallback)
	address := c.GetString("address")
	if data, err := srv.GithubCallback(address, githubCallback, githubCallback.Replace); err != nil {
		OkWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(data, c)
	}
}

// UnbindSocial 解绑
func UnbindSocial(c *gin.Context) {
	var r request.UnbindRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if address == "" {
		Fail(c)
		return
	}
	if err := srv.UnbindSocial(address, r.Type); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

// BindSocialResult 查询绑定结果
func BindSocialResult(c *gin.Context) {
	var r request.BindSocialResultRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if address == "" {
		Fail(c)
		return
	}
	if data, err := srv.BindSocialResult(address, r.Type); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(data, c)
	}
}

// ConfirmBindChange 确认绑定变更
func ConfirmBindChange(c *gin.Context) {
	var r request.ConfirmBindChangeRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if address == "" {
		Fail(c)
		return
	}
	if err := srv.ConfirmBindChange(address, r.Type); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}
