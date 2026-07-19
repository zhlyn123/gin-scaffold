package logger

import (
	"gin-scaffold/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//zap 实列创建

func NewLogger() *zap.Logger {
	cfg := config.GetConfig()

	level := parseLevel(
		cfg.Log.Level,
	)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "time",

		LevelKey: "level",

		CallerKey: "caller",

		MessageKey: "msg",

		EncodeTime: zapcore.ISO8601TimeEncoder,

		EncodeLevel: zapcore.CapitalLevelEncoder,

		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(NewWriter(cfg.Log)),
		level,
	)

	return zap.New(
		core,
		zap.AddCaller(),
	)
}

func parseLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
