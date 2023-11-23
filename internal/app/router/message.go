package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	messageRouter := Router.Group("message").Use(middleware.Auth())
	{
		messageRouter.GET("getUnreadMessage", v1.GetUnreadMessage) // 获取未读消息
		messageRouter.POST("readMessage", v1.ReadMessage)          // 阅读消息
	}
}
