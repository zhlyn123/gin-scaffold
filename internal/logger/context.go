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

func FromContext(c *gin.Context) *zap.Logger {
	requestID := GetRequestID(c)

	// return value.(*zap.Logger)
	return Log.With(
		zap.String("request_id", requestID),
	)
}
