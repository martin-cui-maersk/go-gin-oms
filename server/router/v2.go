package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppRouteV2 struct {
	prefix string // 结构体字段示例
}

// InitOmsAppRouter V2路由
func (s *AppRouteV2) InitOmsAppRouter(r *gin.Engine) {
	omsAppRouterGroup := r.Group(s.prefix)

	// test
	omsAppRouterGroup.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "v2 test"})
	})
}
