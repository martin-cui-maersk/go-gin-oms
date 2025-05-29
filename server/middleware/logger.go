package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LogOptions 定义日志中间件的选项
type LogOptions struct {
	SkipPaths        []string // 不记录日志的路径
	SkipBodyMethods  []string // 不记录请求体的方法（如GET）
	MaxRequestBody   int      // 请求体最大记录长度
	MaxResponseBody  int      // 响应体最大记录长度
	RequestHeaders   []string // 需要记录的请求头
	SkipRecordHeader []string // 需要跳过的请求头
}

// defaultLogOptions 默认日志选项
var defaultLogOptions = LogOptions{
	SkipPaths:        []string{"/health", "/metrics"},
	SkipBodyMethods:  []string{"GET", "HEAD", "OPTIONS"},
	MaxRequestBody:   1024,
	MaxResponseBody:  1024,
	RequestHeaders:   []string{"User-Agent", "Referer", "X-Request-Id"},
	SkipRecordHeader: []string{"Authorization", "Cookie"},
}

// LoggerMiddleware 创建日志中间件
func LoggerMiddleware(logger *zap.Logger, opts ...LogOption) gin.HandlerFunc {
	// 合并选项
	options := defaultLogOptions
	for _, opt := range opts {
		opt(&options)
	}

	// 创建跳过路径的map
	skipPaths := make(map[string]bool)
	for _, path := range options.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()
		// 检查是否需要跳过日志记录
		if skipPaths[c.Request.URL.Path] {
			c.Next()
			return
		}

		// 读取请求体
		var requestBody []byte
		var requestBodyStr string
		if !contains(options.SkipBodyMethods, c.Request.Method) && c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
			requestBodyStr = limitString(string(requestBody), options.MaxRequestBody)
		}

		// 记录请求头
		headerFields := make([]zap.Field, 0)
		for _, h := range options.RequestHeaders {
			if val := c.Request.Header.Get(h); val != "" {
				headerFields = append(headerFields, zap.String("header."+h, val))
			}
		}

		// 创建自定义ResponseWriter来捕获响应
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 计算处理时间
		duration := time.Since(start)

		// 构建日志字段
		fields := []zap.Field{
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("request_body", requestBodyStr),
			zap.String("response_body", limitString(blw.body.String(), options.MaxResponseBody)),
		}
		fields = append(fields, headerFields...)

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

// LogOption 日志选项函数类型
type LogOption func(*LogOptions)

// WithSkipPaths 设置跳过日志的路径
func WithSkipPaths(paths []string) LogOption {
	return func(o *LogOptions) {
		o.SkipPaths = paths
	}
}

// WithMaxBodySize 设置请求/响应体最大记录长度
func WithMaxBodySize(reqMax, respMax int) LogOption {
	return func(o *LogOptions) {
		o.MaxRequestBody = reqMax
		o.MaxResponseBody = respMax
	}
}

// bodyLogWriter 自定义ResponseWriter
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// limitString 限制字符串长度
func limitString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "...[truncated]"
}

// contains 检查字符串是否在切片中
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
