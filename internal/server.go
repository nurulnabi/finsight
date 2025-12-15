package internal

import (
	"os"

	"github.com/gin-gonic/gin"
	middlewares "github.com/nurulnabi/go-finsight/internal/middlewares"
)

type Server interface {
	Init()
}

type WebServer struct {
	Router *gin.Engine
}

func (web *WebServer) Init() {
	r := gin.Default()
	web.Router = r
	r.Use(middlewares.RequestLogger())
	r.Run(os.Getenv("PORT"))
}
