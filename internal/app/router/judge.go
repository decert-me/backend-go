package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitJudgeRouter(Router *gin.RouterGroup) {
	runRouter := Router.Group("judge").Use(middleware.Addr())
	{
		runRouter.POST("/run/tryRun", v1.TryRun)
	}
}
