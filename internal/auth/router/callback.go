package router

import (
	v1 "backend-go/internal/auth/api/v1"
	"backend-go/internal/auth/middleware"
	"github.com/gin-gonic/gin"
)

func InitCallbackRouter(Router *gin.RouterGroup) {
	callbackRouter := Router.Group("callback").Use(middleware.Auth())
	{
		callbackRouter.POST("twitter", v1.TwitterCallback)
		callbackRouter.POST("discord", v1.DiscordCallback)
	}
}
