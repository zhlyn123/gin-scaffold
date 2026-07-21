package logger

import (
	"go.uber.org/zap"
)

// 全局日志实例
var Log *zap.Logger

func InitLogger(l *zap.Logger) {
	Log = l
}

func Close() {
	if Log == nil {
		return
	}
	_ = Log.Sync()
}
