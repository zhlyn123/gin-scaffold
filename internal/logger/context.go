package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//从Context获取Logger

const Loggerkey = "logger"

func SetContextLogger(c *gin.Context, logger *zap.Logger) {
	c.Set(Loggerkey, logger)
}

func FromContext(c *gin.Context) *zap.Logger {
	value, exists := c.Get(Loggerkey)
	if !exists {
		return Log
	}

	return value.(*zap.Logger)
}
