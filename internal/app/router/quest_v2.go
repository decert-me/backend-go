package router

import (
	v2 "backend-go/internal/app/api/v2"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitQuestV2Router(Router *gin.RouterGroup) {
	questRouterAuthV2 := Router.Group("quests").Use(middleware.Auth())
	{
		questRouterAuthV2.POST("", v2.AddQuestV2)
		questRouterAuthV2.PUT("", v2.UpdateQuestV2)
	}
}
