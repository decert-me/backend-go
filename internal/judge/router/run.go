package router

import (
	v1 "backend-go/internal/judge/api/v1"
	"backend-go/internal/judge/middleware"
	"github.com/gin-gonic/gin"
)

func InitRunRouter(Router *gin.RouterGroup) {
	runRouter := Router.Group("run").Use(middleware.Addr())
	{
		runRouter.POST("tryRun", v1.TryRun)
		runRouter.POST("tryTestRun", v1.TryTestRun)
	}
}
