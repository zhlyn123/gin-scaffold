package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//从Context获取Logger

const Loggerkey = "logger"

func GetRequestID(c *gin.Context) string {
	value, exists := c.Get(RequestIDKey)
	if !exists {
		return ""
	}

	return value.(string)
}

func SetContextLogger(c *gin.Context, logger *zap.Logger) {
	c.Set(Loggerkey, logger)
}

func FromContext(c *gin.Context) *zap.Logger {
	requestID := GetRequestID(c)

	// value, exists := c.Get(Loggerkey)
	// if !exists {
	// 	return Log
	// }

	// return value.(*zap.Logger)
	return Log.With(
		zap.String("request_id", requestID),
	)
}
