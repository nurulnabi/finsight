package main

import DB "github.com/nurulnabi/go-finsight/internal/database"

func main() {

	var cfg DB.DBConfig = DB.SQLConfig{
		DB_URI: "postgresql://finsight:hAwJQHsgdaBvUjSGm9zlfw@finsight-19064.j77.aws-ap-south-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full",
	}
	arr := []DB.DBConfig{cfg}

	dbManager := &DB.DatabaseClientManager{
		DBClientsMap: make(map[string]DB.Database),
	}
	dbManager.Init(arr)
}
