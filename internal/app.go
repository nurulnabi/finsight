package internal

import (
	"os"

	"github.com/joho/godotenv"
	DB "github.com/nurulnabi/go-finsight/internal/database"
)

type App struct {
	Name      string
	DbManager *DB.DatabaseClientManager
}

func (app *App) Load() error {
	godotenv.Load(".env")
	var cfg DB.DBConfig = DB.SQLConfig{
		DB_URI: os.Getenv("DB_URI"),
	}
	arr := []DB.DBConfig{cfg}

	app.DbManager = &DB.DatabaseClientManager{
		DBClientsMap: make(map[string]DB.Database),
	}
	app.DbManager.Init(arr)
	return nil
}
