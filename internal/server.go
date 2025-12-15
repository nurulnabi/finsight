package internal

import (
	"os"

	"github.com/gin-gonic/gin"
	middlewares "github.com/nurulnabi/go-finsight/internal/middlewares"
	Routers "github.com/nurulnabi/go-finsight/internal/router"
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
	//register the middlewares
	r.Use(middlewares.RequestLogger())

	//load all other routes
	Routers.Setup(r)

	//start the server
	r.Run(os.Getenv("PORT"))
}
