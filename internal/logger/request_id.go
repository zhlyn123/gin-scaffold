package logger

import (
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(
			"X-Request-ID",
		)

		if requestID == "" {

			requestID = uuid.New().String()

		}

		// 写入响应头

		c.Header(
			"X-Request-ID",
			requestID,
		)

		// 保存到context

		c.Set(
			RequestIDKey,
			requestID,
		)

		c.Next()
	}
}
