package core

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")

	Host := "0.0.0.0"
	if global.CONFIG.System.Env == "public" {
		Host = "127.0.0.1"
	}
	address := fmt.Sprintf("%s:%d", Host, global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 nft-collect
	当前版本:V0.0.1
	默认地址:http://127.0.0.1:%d/
`, global.CONFIG.System.Addr)
	global.LOG.Error(s.ListenAndServe().Error())
}
