package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("users").Use(middleware.Addr())
	userRouterAuth := Router.Group("users").Use(middleware.Auth())
	{
		userRouter.GET("getLoginMessage", v1.GetLoginMessage)
		userRouter.POST("authLoginSign", v1.AuthLoginSign)
		userRouter.GET("/:address", v1.GetUserInfo)
		userRouter.GET("/challenge/:address", v1.GetUserChallengeList)
		userRouter.GET("/quests/:address", v1.GetUserQuestList)
		userRouter.GET("/hasCreateOpenQuestPerm", v1.HasCreateOpenQuestPerm)
	}
	{
		userRouterAuth.GET("discord", v1.GetDiscordInfo)
		userRouterAuth.GET("twitter", v1.GetTwitterInfo)
		userRouterAuth.PUT("/:address", v1.UpdateUserInfo)
		userRouterAuth.POST("/avatar", v1.UpdateAvatar)
		userRouterAuth.GET("/hasBindSocialAccount", v1.HasBindSocialAccount)
	}
}
