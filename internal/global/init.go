package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.SugaredLogger
var LogLevel = zap.NewAtomicLevel()
var DB *gorm.DB
