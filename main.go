package main

import (
	"go-gin-oms/server/models"
	"go-gin-oms/server/routes"
	"os"
)

func init() {
	// 初始化DB连接
	models.ConnectDB()
}

func main() {
	// 启动路由
	r := routes.Routes()
	r.Run("0.0.0.0:" + os.Getenv("HTTP_PORT"))
}
