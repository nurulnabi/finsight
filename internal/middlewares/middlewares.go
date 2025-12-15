package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	Utils "github.com/nurulnabi/go-finsight/internal/config"
)

func HealthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Utils.SendResponse(ctx, map[string]any{
			"data": "ping successful",
		})
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// before request
		c.Next()

		// after request
		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf(
			"method=%s path=%s status=%d latency=%s",
			c.Request.Method,
			c.Request.URL.Path,
			status,
			latency,
		)
	}
}
