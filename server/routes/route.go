package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1User "go-gin-oms/server/api/v1/user"
	"go-gin-oms/server/global"
	"go-gin-oms/server/middlewares"
	"net/http"
	"runtime/debug"
	"time"
)

func Routes() *gin.Engine {
	r := gin.Default()

	// 日志中间件
	r.Use(middlewares.LoggerMiddleware(global.ZapLogger))

	//处理找不到路由
	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)

	// 处理发生异常
	r.Use(Recover)

	api := r.Group("/api")
	{
		omsApp := api.Group("/oms-app/v1")
		{
			// public路由
			public := omsApp.Group("/user")
			{
				public.POST("/login", v1User.Login)
			}

			// protected路由
			protected := omsApp.Group("/")
			{
				protected.Use(middlewares.JwtAuthMiddleware())
				userGroup := protected.Group("/user")
				{
					userGroup.GET("/info", v1User.CurrentUserInfo)
					userGroup.GET("/permission-code", v1User.CurrentUserInfo)
					userGroup.GET("/menu-list", v1User.GetMyMenuList)
				}

			}

		}
	}
	return r
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
			//return
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "System busy, please try again later!"})
		}
	}()
	//继续后续接口调用
	c.Next()
}
