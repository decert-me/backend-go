package router

import (
	v1 "backend-go/internal/auth/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("authorization")
	{
		authRouter.GET("discord", v1.DiscordAuthorizationURL) // 获取discord授权地址
		authRouter.GET("wechat", v1.GetWechatQrcode)          // 获取关注二维码
		authRouter.GET("github", v1.GithubAuthorizationURL)   // 获取github授权地址
	}
}
