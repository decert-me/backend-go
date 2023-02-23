package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("users")
	userRouterAuth := Router.Group("users").Use(middleware.Auth())
	{
		userRouter.GET("getLoginMessage", v1.GetLoginMessage)
		userRouter.POST("authLoginSign", v1.AuthLoginSign)
	}
	{
		userRouterAuth.GET("discord", v1.GetDiscordInfo)
	}
}
