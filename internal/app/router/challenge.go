package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitChallengeRouter(Router *gin.RouterGroup) {
	challengeAuthRouter := Router.Group("challenge").Use(middleware.Addr())
	{
		challengeAuthRouter.POST("", v1.CreateChallengeLog)
	}
}
