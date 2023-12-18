package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/config"
	"backend-go/internal/app/middleware"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
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
	// sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           c.Sentry.Dsn,
		EnableTracing: c.Sentry.EnableTracing,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: c.Sentry.TracesSampleRate,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	} else {
		Router.Use(sentrygin.New(sentrygin.Options{}))
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
	v1Group.Use(middleware.I18n())
	{
		InitQuestRouter(v1Group)
		InitUserRouter(v1Group)
		InitBadgeRouter(v1Group)
		InitTransactionRouter(v1Group)
		InitChallengeRouter(v1Group)
		InitIPFSRouter(v1Group)
		InitJudgeRouter(v1Group)
		InitTutorialRouter(v1Group)
		InitShareRouter(v1Group)
		InitCollectionRouter(v1Group)
		InitMessageRouter(v1Group)
		InitSocialRouter(v1Group)
		InitOpenQuestRouter(v1Group)
	}
	// meta
	Router.GET("/quests/:id", v1.HandleMetaRequest)
	Router.GET("/claim/:id", v1.HandleMetaRequest)
	Router.GET("/collection/:id", v1.HandleCollectionMetaRequest)

	fmt.Println("router register success")
	return Router
}
