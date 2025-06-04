package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-oms/server/api/v1/system"
	"go-gin-oms/server/api/v1/user"
	"go-gin-oms/server/middleware"
)

type AppRouteV1 struct {
	prefix string // 结构体字段示例
}

// InitOmsAppRouter V1路由
func (s *AppRouteV1) InitOmsAppRouter(r *gin.Engine) {
	omsAppRouterGroup := r.Group(s.prefix)

	// 登录路由 public权限
	omsAppRouterGroup.POST("/user/login", user.Login)

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
