package router

import (
	v1 "backend-go/internal/judge/api/v1"
	"backend-go/internal/judge/middleware"
	"github.com/gin-gonic/gin"
)

func InitBuildRouter(Router *gin.RouterGroup) {
	questRouter := Router.Group("build").Use(middleware.Addr())
	{
		questRouter.POST("", v1.BuildSolidity)
	}
}
