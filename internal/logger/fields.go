package logger

import "go.uber.org/zap"

// 通用日志字段

func RequestField(id string) zap.Field {
	return zap.String("request_id", id)
}