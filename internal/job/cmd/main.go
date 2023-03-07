package main

import (
	"backend-go/internal/job/initialize"
	"backend-go/internal/job/service"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		log.Infov("backend-job get a signal", zap.String("signal", si.String()))
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("backend-job exit")
			s.Close()
			time.Sleep(2 * time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
