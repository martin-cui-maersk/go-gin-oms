package core

import (
	"fmt"
	"go-gin-oms/server/config"
	"go.uber.org/zap"
	"time"
)

// InitLogger 初始化日志
func InitLogger() (*zap.Logger, *zap.Logger) {
	// 初始化Zap日志记录器
	accessLogFileName := fmt.Sprintf("logs/access_log/%s.log", time.Now().Format("2006-01-02"))
	accessLogger, err := config.NewLogger(config.LogConfig{
		Filename:   accessLogFileName,
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
		Level:      "info",
		Console:    true,
	})
	if err != nil {
		panic(err)
	}
	defer accessLogger.Sync()

	// 初始化应用日志记录器
	appLogFileName := fmt.Sprintf("logs/app_log/%s.log", time.Now().Format("2006-01-02"))
	appLogger, err := config.NewLogger(config.LogConfig{
		Filename:   appLogFileName,
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
		Level:      "debug",
		Console:    true,
	})
	if err != nil {
		panic(err)
	}
	defer appLogger.Sync()

	return accessLogger, appLogger
}
