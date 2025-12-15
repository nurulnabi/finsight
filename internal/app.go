package internal

import (
	"os"

	DB "github.com/nurulnabi/go-finsight/internal/database"
)

type App struct {
	Name      string
	AppType   string
	DbManager *DB.DatabaseClientManager
	server    Server
}

func (app *App) Load() error {
	if app.AppType == "" {
		app.AppType = "WEB_SERVER"
	}

	var cfg DB.DBConfig = DB.SQLConfig{
		DB_URI: os.Getenv("DB_URI"),
	}
	arr := []DB.DBConfig{cfg}

	app.DbManager = &DB.DatabaseClientManager{
		DBClientsMap: make(map[string]DB.Database),
	}
	app.DbManager.Init(arr)

	web := WebServer{}
	app.server = &web
	web.Init()
	return nil
}
