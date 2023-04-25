package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitBadgeRouter(Router *gin.RouterGroup) {
	badgeRouter := Router.Group("badge").Use(middleware.Auth())
	{
		badgeRouter.POST("submitClaimTweet", v1.SubmitClaimTweet)
		badgeRouter.POST("claim", v1.PermitClaimBadge)
	}
}
