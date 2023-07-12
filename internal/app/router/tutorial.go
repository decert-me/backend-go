package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitTutorialRouter(Router *gin.RouterGroup) {
	tutorialRouter := Router.Group("tutorial").Use(middleware.Addr())
	//tutorialRouterAuth := Router.Group("tutorial").Use(middleware.Auth())
	{
		tutorialRouter.POST("progress", v1.GetProgress)
		tutorialRouter.PUT("progress", v1.UpdateProgress)
	}
}
