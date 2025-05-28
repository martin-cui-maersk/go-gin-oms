package main

import (
	"go-gin-oms/server/core"
	"go-gin-oms/server/routes"
	"os"
)

func init() {
	// 初始化Zap日志记录器

	// 初始化数据库
	core.ConnectDB()
}

func main() {
	// 启动路由
	r := routes.Routes()
	r.Run("0.0.0.0:" + os.Getenv("HTTP_PORT"))
}
