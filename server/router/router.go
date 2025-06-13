package router

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/martin-cui-maersk/go-gin-oms/global"
	"github.com/martin-cui-maersk/go-gin-oms/middleware"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
	"time"
)

func InitRouter() {
	// 创建Gin引擎
	r := gin.New()

	// 处理发生异常
	r.Use(Recover)

	// 处理找不到路由
	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)

	// 日志中间件
	r.Use(middleware.LoggerMiddleware(global.AccessLogger,
		middleware.WithSkipPaths([]string{"/health", "/favicon.ico"}),
		middleware.WithMaxBodySize(2048, 2048),
	))
	{
		// 心跳检测路由
		r.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})
	}
	{
		// v1版本路由
		v1Route := &AppRouteV1{prefix: "/api/oms-app/v1"} // 创建指针实例
		v1Route.InitOmsAppRouter(r)                       // 调用方法
	}
	{
		// v2版本路由
		v2Route := &AppRouteV2{prefix: "/api/oms-app/v2"} // 创建指针实例
		v2Route.InitOmsAppRouter(r)                       // 调用方法
	}

	//err := r.Run(fmt.Sprintf(":%s", global.Server.Port))
	//if err != nil {
	//	global.AppLogger.Fatal("Failed to start server", zap.Error(err))
	//}

	// 使用 endless 监听端口
	err := endless.ListenAndServe(":"+global.Server.Port, r)
	if err != nil {
		global.AppLogger.Fatal("Failed to start server", zap.Error(err))
	}
}

// HandleNotFound 404 找不到路径时的处理
func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 404, "msg": "Page not found"})
}

// Recover 500 内部发生异常时的处理
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//打印错误堆栈信息
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("当前时间:", timeStr)
			fmt.Println("当前访问path:", c.FullPath())
			fmt.Println("当前完整地址:", c.Request.URL.String())
			fmt.Println("当前协议:", c.Request.Proto)
			//fmt.Println("当前get参数:", result.GetAllGetParams(c))
			//fmt.Println("当前post参数:", result.GetAllPostParams(c))
			fmt.Println("当前访问方法:", c.Request.Method)
			fmt.Println("当前访问Host:", c.Request.Host)
			fmt.Println("当前IP:", c.ClientIP())
			fmt.Println("当前浏览器:", c.Request.UserAgent())
			fmt.Println("发生异常:", err)
			//result.Logger.Errorf("stack: %v",string(debug.Stack()))
			debug.PrintStack()
			global.AppLogger.Fatal(fmt.Sprintf("stack: %v", err))
			//return
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "System busy, please try again later!"})
		}
	}()
	// 继续后续接口调用
	c.Next()
}
