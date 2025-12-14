package main

import DB "github.com/nurulnabi/go-finsight/internal/database"

func main() {

	var cfg DB.DBConfig = DB.SQLConfig{
		DB_URI: "",
	}
	arr := []DB.DBConfig{cfg}

	dbManager := &DB.DatabaseClientManager{
		DBClientsMap: make(map[string]DB.Database),
	}
	dbManager.Init(arr)
}
