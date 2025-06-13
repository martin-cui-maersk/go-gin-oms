package router

import (
	"github.com/gin-gonic/gin"
	"github.com/martin-cui-maersk/go-gin-oms/api/v1/system"
	"github.com/martin-cui-maersk/go-gin-oms/api/v1/user"
	"github.com/martin-cui-maersk/go-gin-oms/global"
	"github.com/martin-cui-maersk/go-gin-oms/middleware"
)

type AppRouteV1 struct {
	prefix string // 结构体字段示例
}

// InitOmsAppRouter V1路由
func (s *AppRouteV1) InitOmsAppRouter(r *gin.Engine) {
	// prefix
	omsAppRouterGroup := r.Group(s.prefix)

	// 登录路由 public权限
	omsAppRouterGroup.POST("/user/login", user.Login)
	omsAppRouterGroup.GET("/version", func(c *gin.Context) {
		c.JSON(200, global.Server)
	})

	// jwt中间件
	omsAppRouterGroup.Use(middleware.JwtAuthMiddleware())

	// 用户路由
	userGroup := omsAppRouterGroup.Group("/user")
	{
		userGroup.GET("/info", user.CurrentUserInfo)
		userGroup.GET("/permission-code", user.GetPermissionCode)
		userGroup.GET("/menu-list", user.GetMyMenuList)
	}

	// 系统路由
	systemGroup := omsAppRouterGroup.Group("/system")
	{
		// user
		systemGroup.GET("/user-list", system.GetUserList)
		systemGroup.GET("/role-ids", system.GetRoleSelect)

		// role
		systemGroup.GET("/role-list", system.GetRoleList)
		systemGroup.POST("/add-role", system.AddRole)
		systemGroup.POST("/update-role", system.UpdateRole)
		systemGroup.POST("/set-role-status", system.SetRoleStatus)

		// menu
		systemGroup.GET("/menu-list", system.GetMenuList)
		systemGroup.POST("/add-menu", system.AddMenu)
		systemGroup.POST("/update-menu", system.UpdateMenu)
		systemGroup.POST("/delete-menu", system.DeleteMenu)

		// store TODO::待开发
		// merchant TODO::待开发
	}
}
