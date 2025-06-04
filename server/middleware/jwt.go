package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-oms/server/utils/result"
	"go-gin-oms/server/utils/token"
)

// JwtAuthMiddleware JWT中间件
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.CheckTokenValid(c)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
			//c.Abort()
			result.Response().SetCode(600).SetMsg("Unauthorized").SetData(nil).Build(c)
			return
		}
		c.Next()
	}
}
