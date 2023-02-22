package main

import (
	"backend-go/internal/app/core"
	"backend-go/internal/app/global"
	"backend-go/internal/app/initialize"
	"backend-go/internal/app/timer"
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
	// 初始化合约信息
	initialize.InitContract()
	// 初始化缓存
	initialize.Redis()
	// 定时任务
	timer.Timer()
	core.RunWindowsServer()
}
