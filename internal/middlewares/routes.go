package middlewares

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/healthcheck", HealthCheck())
}
