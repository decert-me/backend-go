package main

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/initialize"
	"backend-go/internal/app/router"
	"backend-go/internal/app/service"
	"backend-go/pkg/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化Viper
	config := initialize.Viper()
	// 初始化日志框架
	log.Init(config.Log)
	// 初始化service
	s := service.New(config)
	// 初始化api
	v1.Init(s)
	// 初始化router
	router.New(config)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		log.Infov("backend-go get a signal", zap.String("signal", si.String()))
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("backend-go exit")
			s.Close()
			time.Sleep(2 * time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
