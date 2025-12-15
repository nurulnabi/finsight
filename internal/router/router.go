package router

import (
	"github.com/gin-gonic/gin"
	Middlewares "github.com/nurulnabi/go-finsight/internal/middlewares"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api/v1")

	Middlewares.RegisterRoutes(api)
}
