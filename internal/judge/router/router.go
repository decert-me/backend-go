package router

import (
	v1 "backend-go/internal/judge/api/v1"
	"backend-go/internal/judge/config"
	"backend-go/internal/judge/middleware"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func New(c *config.Config) {
	middleware.Init(c)
	Router := Routers(c)
	Host := "0.0.0.0"
	if c.System.Env == "public" {
		Host = "127.0.0.1"
	}
	address := fmt.Sprintf("%s:%d", Host, c.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	log.Printf("server run success on %s", address)

	fmt.Printf(`
	欢迎使用 backend-go
	当前版本:V0.0.1
	默认地址:http://127.0.0.1:%d/`, c.System.Addr)
	fmt.Println()
	s.ListenAndServe()
}

// Routers 初始化总路由
func Routers(c *config.Config) *gin.Engine {
	var Router *gin.Engine
	// 开发环境打开日志 && 打开pprof
	if c.System.Env == "develop" {
		Router = gin.Default()
		pprof.Register(Router) // 性能
	} else {
		Router = gin.New()
		Router.Use(gin.Recovery())
	}
	Router.Use(middleware.Cors()) // 放行跨域请求
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			v1.Ping(c)
		})
	}
	v1Group := Router.Group("v1")
	{
		InitBuildRouter(v1Group)
		InitRPCRouter(v1Group)
		InitCastRouter(v1Group)
		InitRunRouter(v1Group)
	}

	fmt.Println("router register success")
	return Router
}
