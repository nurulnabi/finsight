package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	BaseApp "github.com/nurulnabi/go-finsight/internal"
)

func main() {
	godotenv.Load(".env")

	app := BaseApp.App{
		Name:    os.Getenv("APP_NAME"),
		AppType: os.Getenv("APP_TYPE"),
	}

	err := app.Load()
	if err != nil {
		fmt.Println("Error while Loading the App")
	}
}
