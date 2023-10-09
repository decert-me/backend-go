package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitBadgeRouter(Router *gin.RouterGroup) {
	badgeRouter := Router.Group("badge").Use(middleware.Auth())
	badgeAddrRouter := Router.Group("badge").Use(middleware.Addr())
	{
		badgeRouter.POST("submitClaimTweet", v1.SubmitClaimTweet)
		badgeRouter.POST("submitClaimShare", v1.SubmitClaimShare)
		badgeRouter.POST("claim", v1.PermitClaimBadge)
		badgeRouter.PUT("uri", v1.UpdateBadgeURI)
	}
	{
		badgeAddrRouter.GET("hasClaimed/:tokenID", v1.HasClaimed)
	}
}
