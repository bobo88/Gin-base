package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		gin.DefaultWriter.Write([]byte(
			"[GIN] " + time.Now().Format("2006-01-02 15:04:05") +
				" | " + c.Request.Method +
				" | " + path +
				" | " + string(rune(statusCode)) +
				" | " + latency.String() + "\n",
		))
	}
}