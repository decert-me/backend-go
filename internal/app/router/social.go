package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialRouter(Router *gin.RouterGroup) {
	routerAuth := Router.Group("social").Use(middleware.Auth()).Use(middleware.I18n())
	router := Router.Group("social")
	{
		routerAuth.GET("/getWechatQrcode", v1.GetWechatQrcode)        // 获取关注二维码
		routerAuth.POST("/discordBindAddress", v1.DiscordBindAddress) // Discord 回调绑定
	}
	{
		router.POST("/wechatBindAddress", v1.WechatBindAddress)               // 微信绑定
		router.GET("/getDiscordAuthorizationURL", v1.DiscordAuthorizationURL) // 获取 Discord 授权链接
	}
	{
		routerAuth.POST("/getEmailBindCode", v1.GetEmailBindCode) // 获取邮箱绑定验证码
		routerAuth.POST("/emailBindAddress", v1.EmailBindAddress) // 绑定邮箱
	}
	{
		routerAuth.GET("/getGithubAuthorizationURL", v1.GithubAuthorizationURL) // 获取 Github 授权链接
		routerAuth.POST("/githubBindAddress", v1.GithubBindAddress)             // Github 回调绑定
	}
	{
		routerAuth.POST("/bindSocialResult", v1.BindSocialResult)   // 查询绑定结果
		routerAuth.POST("/confirmBindChange", v1.ConfirmBindChange) // 确认绑定
		routerAuth.POST("/unbindSocial", v1.UnbindSocial)           // 解绑
	}
}
