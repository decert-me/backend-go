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
		collectionRouter.GET("challengeUsers", v1.GetCollectionChallengeUser)
		collectionRouter.GET("", v1.GetCollectionQuest)
	}
	{
		questRouterAuth.POST("claim", v1.CollectionClaim)
	}
}
