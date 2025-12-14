package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	Errors "github.com/nurulnabi/go-finsight/internal/errors"
)

type Database interface {
	// Database this interface must be implemented by all the db providers check sql.go
	Connect(cfg DBConfig) error
	HealthCheck(ctx context.Context) error
	Close() error
}

type DatabaseClientManager struct {
	DBClientsMap map[string]Database
}

// init initialize all the kinds of databases required for this application.
//
// Parameters:
//   - cfgs: array of database-specific configuration, can be of type SQLConfig, MongoConfig.
func (dc *DatabaseClientManager) Init(cfgs []DBConfig) {
	for _, dbConfig := range cfgs {
		switch dbConfig.GetDriver() {
		case "sql":
			dc.initSQL(dbConfig)
		default:
			return
		}
	}
}

// initSQL establishes a connection to the database.
//
// Parameters:
//   - cfg: provides database-specific configuration.
//
// Returns:
//   - *Errors.AppError: non-nil if the connection fails.
func (dc *DatabaseClientManager) initSQL(cfg DBConfig) error {
	var sqlConfig SQLConfig
	b, _ := json.Marshal(cfg)
	err := json.Unmarshal(b, &sqlConfig)
	if err != nil {
		fmt.Println(err)
		return Errors.AppError{}.DbConnectionError("Invalid SQL config")
	}
	var errr error
	var sqlDB Database = &SqlDatabase{}
	fmt.Println("asdfads", sqlDB)
	log.Printf("ptr address: %p", sqlDB)
	errr = sqlDB.Connect(sqlConfig)
	if errr != nil {
		fmt.Println(errr.Error())
		return errr
	}
	log.Printf("ptr address: %p", sqlDB)
	fmt.Println("asdfad", sqlDB)
	dc.DBClientsMap[sqlConfig.NAME] = sqlDB
	return nil
}

func (dc *DatabaseClientManager) GetDBClient(name string) *Database {
	client, ok := dc.DBClientsMap[name]
	if !ok {
		err := Errors.AppError{}.NotFoundError("DatabaseClientManager.GetDBClient client not found- " + name)
		err.LogError()
	}
	return &client
}
