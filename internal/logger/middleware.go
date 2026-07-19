package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//Gin请求日志

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()
		cost := time.Since(start)
		Log.Info(
			"http request",

			zap.String(
				"method",
				method,
			),

			zap.String(
				"path",
				path,
			),

			zap.Int(
				"status",
				c.Writer.Status(),
			),

			zap.Duration(
				"latency",
				cost,
			),
		)
	}
}
