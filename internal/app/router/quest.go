package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitQuestRouter(Router *gin.RouterGroup) {
	questRouter := Router.Group("quests").Use(middleware.Addr())
	questRouterAuth := Router.Group("quests").Use(middleware.Auth())
	{
		questRouter.GET("", v1.GetQuestList)
		questRouter.GET("/:id", v1.GetQuest)
		questRouter.GET("/:id/challengeUsers", v1.GetQuestChallengeUser)
		questRouter.GET("/:id/getQuestFlashRank", v1.GetQuestFlashRank) // 获取闪电榜

	}
	{
		questRouterAuth.POST("", v1.AddQuest)
		questRouterAuth.PUT("", v1.UpdateQuest)
		questRouterAuth.PUT("/recommend", v1.UpdateRecommend)
	}
}
