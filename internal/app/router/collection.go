package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitCollectionRouter(Router *gin.RouterGroup) {
	collectionRouter := Router.Group("collection").Use(middleware.Addr())
	questRouterAuth := Router.Group("collection").Use(middleware.Auth())
	{
		collectionRouter.GET("challengeUsers", v1.GetCollectionChallengeUser)     // 获取挑战合辑用户
		collectionRouter.GET("", v1.GetCollectionQuest)                           // 获取合辑内挑战
		collectionRouter.GET("checkQuestInCollection", v1.CheckQuestInCollection) // 查询挑战是否在合辑内
	}
	{
		questRouterAuth.POST("claim", v1.CollectionClaim) // 领取合辑奖励
	}
}
