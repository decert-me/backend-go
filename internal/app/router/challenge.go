package router

import (
	v1 "backend-go/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitChallengeRouter(Router *gin.RouterGroup) {
	challengeAuthRouter := Router.Group("challenge")
	{
		challengeAuthRouter.POST("", v1.CreateChallengeLog)
	}
}
