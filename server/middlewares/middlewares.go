package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-gin-oms/server/utils/global"
	"go-gin-oms/server/utils/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.CheckTokenValid(c)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
			//c.Abort()
			global.NewResult().SetCode(401).SetMsg("Unauthorized").SetData(nil).Build(c)
			return
		}
		c.Next()
	}
}
