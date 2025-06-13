package main

import (
	"github.com/martin-cui-maersk/go-gin-oms/core"
	"github.com/martin-cui-maersk/go-gin-oms/global"
	"github.com/martin-cui-maersk/go-gin-oms/router"
)

func init() {
	// 初始化配置
	global.Config = core.Viper()
	// 初始化日志
	global.AccessLogger, global.AppLogger = core.InitLogger()
	// 初始化数据库
	global.DB = core.ConnectDB()
	// 初始化Redis
	// TODO::初始化Redis
}

func main() {
	// 启动路由
	router.InitRouter()
}
