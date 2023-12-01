package router

import (
	v2 "backend-go/internal/app/api/v2"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitCollectionV2Router(Router *gin.RouterGroup) {
	questRouterAuth := Router.Group("collection").Use(middleware.Auth())
	{
		questRouterAuth.POST("claim", v2.CollectionClaimV2) // 领取合辑奖励
	}
}
