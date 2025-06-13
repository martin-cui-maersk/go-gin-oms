package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/martin-cui-maersk/go-gin-oms/utils/result"
	"github.com/martin-cui-maersk/go-gin-oms/utils/token"
)

// JwtAuthMiddleware JWT中间件
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.CheckTokenValid(c)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
			//c.Abort()
			result.Response().SetCode(600).SetMsg("Token expired, please login").SetData(nil).Build(c)
			return
		}
		c.Next()
	}
}
