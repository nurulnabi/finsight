package main

import (
	"os"

	BaseApp "github.com/nurulnabi/go-finsight/internal"
)

func main() {
	app := BaseApp.App{
		Name: os.Getenv("APP_NAME"),
	}

	app.Load()
}
