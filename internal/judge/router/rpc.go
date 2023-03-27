package router

import (
	v1 "backend-go/internal/judge/api/v1"
	"backend-go/internal/judge/middleware"
	"github.com/gin-gonic/gin"
)

func InitRPCRouter(Router *gin.RouterGroup) {
	questRouter := Router.Group("rpc").Use(middleware.Addr())
	{
		questRouter.POST("", v1.HandleProxy)
	}
}
