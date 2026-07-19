package logger

import (
	"gin-scaffold/internal/config"
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

// 文件切割
func NewWriter(cfg config.LogConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
}
