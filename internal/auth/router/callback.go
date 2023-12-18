package router

import (
	v1 "backend-go/internal/auth/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCallbackRouter(Router *gin.RouterGroup) {
	callbackRouter := Router.Group("callback")
	{
		callbackRouter.POST("discord", v1.DiscordCallback)
	}
}
