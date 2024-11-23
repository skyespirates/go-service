package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		Logger.Infow("Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration,
		)
	}
}
