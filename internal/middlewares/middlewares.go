package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
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
