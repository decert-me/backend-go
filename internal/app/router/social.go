package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitSocialRouter(Router *gin.RouterGroup) {
	routerAuth := Router.Group("social").Use(middleware.Auth())
	router := Router.Group("social")
	{
		routerAuth.GET("/getWechatQrcode", v1.GetWechatQrcode)        // 获取关注二维码
		routerAuth.POST("/discordBindAddress", v1.DiscordBindAddress) // Discord 回调绑定
	}
	{
		router.POST("/wechatBindAddress", v1.WechatBindAddress)               // 微信绑定
		router.GET("/getDiscordAuthorizationURL", v1.DiscordAuthorizationURL) // 获取 Discord 授权链接
	}
}
