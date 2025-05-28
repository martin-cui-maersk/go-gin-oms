package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-gin-oms/server/utils/result"
	"go-gin-oms/server/utils/token"
	"go.uber.org/zap"
	"io"
	"time"
)

// JwtAuthMiddleware JWT中间件
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.CheckTokenValid(c)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Unauthorized"})
			//c.Abort()
			result.NewResult().SetCode(401).SetMsg("Unauthorized").SetData(nil).Build(c)
			return
		}
		c.Next()
	}
}

// LoggerMiddleware 日志中间件
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 读取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后需要重新设置Body，因为ReadAll消耗了它
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建一个自定义的ResponseWriter来捕获响应
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// next
		c.Next()

		// 记录请求结束时间
		duration := time.Since(start)

		// 记录请求和响应信息
		fields := []zap.Field{
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("proto", c.Request.Proto),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("request_body", string(requestBody)),
			zap.String("response_body", blw.body.String()),
		}

		// 根据状态码决定日志级别
		status := c.Writer.Status()
		switch {
		case status >= 500:
			logger.Error("Server error", fields...)
		case status >= 400:
			logger.Warn("Client error", fields...)
		default:
			logger.Info("Request processed", fields...)
		}
	}
}

// bodyLogWriter 自定义的ResponseWriter，用于捕获响应体
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
