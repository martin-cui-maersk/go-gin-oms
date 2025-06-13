package core

import (
	"context"
	"fmt"
	"github.com/martin-cui-maersk/go-gin-oms/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Redis 连接Redis
func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Server.Redis.Host, global.Server.Redis.Port),
		Password: global.Server.Redis.Password,
		DB:       global.Server.Redis.Database,
	})
	//defer func(rdb *redis.Client) {
	//	err := rdb.Close()
	//	if err != nil {
	//		fmt.Printf("close redis failed, error is:  %v\n", err)
	//	}
	//}(rdb)
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("redis connect ping failed, %w", err))
	}
	global.AppLogger.Info("redis connect ping response:", zap.String("name", fmt.Sprintf("%s:%d, ping result: %s", global.Server.Redis.Host, global.Server.Redis.Port, pong)))
	return rdb
}
