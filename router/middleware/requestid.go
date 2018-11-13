package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查上行数据的requestID
		requestID := c.Request.Header.Get("X-Request-Id")

		// 通过UUID4创建requestID
		if requestID == "" {
			u4 := uuid.NewV4()
			requestID = u4.String()
		}

		c.Set("X-Request-Id", requestID)

		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
