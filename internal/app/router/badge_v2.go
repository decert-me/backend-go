package router

import (
	v2 "backend-go/internal/app/api/v2"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitBadgeV2Router(Router *gin.RouterGroup) {
	badgeRouter := Router.Group("badge").Use(middleware.Auth())
	{
		badgeRouter.POST("submitClaimShare", v2.SubmitClaimShareV2)
		badgeRouter.POST("generateMintSignature", v2.GenerateMintSignature)
		badgeRouter.POST("confirmUserMint", v2.ConfirmUserMint)
	}
}
