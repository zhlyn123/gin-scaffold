package database

import (
	"context"
	"gin-scaffold/internal/logger"
	"time"

	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	Level         gormLogger.LogLevel
	SlowThreshold time.Duration
}

func NewGormLogger(
	level gormLogger.LogLevel,
	slow time.Duration,
) *GormLogger {

	return &GormLogger{
		Level:         level,
		SlowThreshold: slow,
	}

}

func (l *GormLogger) LogMode(
	level gormLogger.LogLevel,
) gormLogger.Interface {
	newLogger := *l
	newLogger.Level = level

	return &newLogger

}

func (l *GormLogger) Info(
	ctx context.Context,
	msg string,
	args ...interface{},
) {

	logger.Log.Info(
		msg,
	)

}

func (l *GormLogger) Warn(
	ctx context.Context,
	msg string,
	args ...interface{},
) {

	logger.Log.Warn(
		msg,
	)

}

func (l *GormLogger) Error(
	ctx context.Context,
	msg string,
	args ...interface{},
) {

	logger.Log.Error(
		msg,
	)

}

func (l *GormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (string, int64),
	err error,
) {

	requestID := getRequestID(ctx)

	elapsed := time.Since(begin)

	sql, rows := fc()

	fields := []zap.Field{
		zap.String(
			"request_id",
			requestID,
		),

		zap.String(
			"sql",
			sql,
		),

		zap.Int64(
			"rows",
			rows,
		),

		zap.Duration(
			"duration",
			elapsed,
		),
	}

	if err != nil {

		logger.Log.Error(
			"database error",

			append(
				fields,
				zap.Error(err),
			)...,
		)

		return

	}

	if elapsed > l.SlowThreshold {

		logger.Log.Warn(
			"database query slow",
			fields...,
		)

	}

	logger.Log.Debug(
		"database query",
		fields...,
	)

}

func getRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(logger.RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}

func ConvertGormLogLevel(
	level string,
) gormLogger.LogLevel {
	switch level {

	case "debug":

		return gormLogger.Info

	case "error":

		return gormLogger.Error

	default:

		return gormLogger.Warn

	}
}
