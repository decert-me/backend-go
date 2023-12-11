package router

import (
	v1 "backend-go/internal/auth/api/v1"
	"backend-go/internal/auth/middleware"
	"github.com/gin-gonic/gin"
)

func InitWechatRouter(Router *gin.RouterGroup) {
	wechatRouter := Router.Group("")
	wechatRouterWithAuth := Router.Group("wechat").Use(middleware.Auth())
	{
		wechatRouter.Any("/", v1.WechatService) // 微信服务器验证
	}
	{
		wechatRouterWithAuth.GET("/getWechatQrcode", v1.GetWechatQrcode) // 获取关注二维码
	}
}
