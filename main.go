package main

import (
	"go-gin-oms/server/core"
	"go-gin-oms/server/global"
	"go-gin-oms/server/routes"
	"go.uber.org/zap"
	"os"
)

func init() {
	// 初始化日志
	global.AccessLogger, global.AppLogger = core.InitLogger()
	// 初始化数据库
	global.DB = core.ConnectDB()
}

func main() {
	// 启动路由
	r := routes.Routes()
	err := r.Run("0.0.0.0:" + os.Getenv("HTTP_PORT"))
	if err != nil {
		global.AppLogger.Fatal("Failed to start server", zap.Error(err))
		return
	}
}
