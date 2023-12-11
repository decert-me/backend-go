package router

import (
	v1 "backend-go/internal/auth/api/v1"
	"backend-go/internal/auth/middleware"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("authorization").Use(middleware.Auth())
	{
		authRouter.GET("twitter", v1.TwitterAuthorizationURL)
		authRouter.GET("discord", v1.DiscordAuthorizationURL)
	}
}
