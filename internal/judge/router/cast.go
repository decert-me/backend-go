package router

import (
	v1 "backend-go/internal/judge/api/v1"
	"backend-go/internal/judge/middleware"
	"github.com/gin-gonic/gin"
)

func InitCastRouter(Router *gin.RouterGroup) {
	questRouter := Router.Group("cast").Use(middleware.Addr())
	{
		questRouter.POST("call", v1.CastCall)
		questRouter.POST("send", v1.CastSend)
	}
}
