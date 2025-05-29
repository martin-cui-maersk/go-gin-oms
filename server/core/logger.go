package core

import (
	"go-gin-oms/server/config"
	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, *zap.Logger) {
	// 初始化Zap日志记录器
	accessLogger, err := config.NewLogger(config.LogConfig{
		Filename:   "logs/access.log",
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
	appLogger, err := config.NewLogger(config.LogConfig{
		Filename:   "logs/app.log",
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
