package logger

import (
	"gin-scaffold/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//zap 实列创建

func NewLogger() *zap.Logger {
	//获取全局配置
	cfg := config.GetConfig()

	//定义记录日志级别
	level := parseLevel(
		cfg.Log.Level,
	)

	//定义输出格式模板
	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "time",

		LevelKey: "level",

		CallerKey: "caller",

		MessageKey: "msg",

		EncodeTime: zapcore.ISO8601TimeEncoder,

		EncodeLevel: zapcore.CapitalLevelEncoder,

		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	//创建 JSON 编码器
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	//创建 输出核心
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
