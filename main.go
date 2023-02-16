package main

import (
	"backend-go/internal/app/core"
	"backend-go/internal/app/global"
	"backend-go/internal/app/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化Viper
	core.Viper()
	// 初始化zap日志库
	global.LOG = core.Zap()
	// 注册全局logger
	zap.ReplaceGlobals(global.LOG)
	// 初始化数据库
	initialize.InitCommonDB()
	core.RunWindowsServer()
}
