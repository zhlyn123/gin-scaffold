package logger

import (
	"go.uber.org/zap"
)

// 全局日志实例
var Log *zap.Logger

func InitLogger(logger *zap.Logger) {
	Log = logger
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

