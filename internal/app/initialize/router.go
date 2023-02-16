package initialize

import (
	"backend-go/internal/app/global"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router *gin.Engine
	// 开发环境打开日志 && 打开pprof
	if global.CONFIG.System.Env == "develop" {
		Router = gin.Default()
		pprof.Register(Router) // 性能
	} else {
		Router = gin.New()
		Router.Use(gin.Recovery())
	}
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	v1Group := Router.Group("v1")
	{
		_ = v1Group
		//router.InitAccountRouter(v1Group)
	}

	global.LOG.Info("router register success")
	return Router
}
