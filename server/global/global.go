package global

import (
	"github.com/martin-cui-maersk/go-gin-oms/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Redis        *redis.Client
	DB           *gorm.DB
	AccessLogger *zap.Logger
	AppLogger    *zap.Logger
	Config       *viper.Viper
	Server       config.Server
)
