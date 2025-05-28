package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ZapLogger *zap.Logger
var DB *gorm.DB
