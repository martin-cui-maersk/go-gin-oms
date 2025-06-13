package global

import (
	"github.com/martin-cui-maersk/go-gin-oms/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	AccessLogger *zap.Logger
	AppLogger    *zap.Logger
	Config       *viper.Viper
	Server       config.Server
)
