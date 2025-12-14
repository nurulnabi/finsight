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
	Connect(cfg DBConfig) *Errors.AppError
	HealthCheck(ctx context.Context) *Errors.AppError
	Close() *Errors.AppError
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
func (dc *DatabaseClientManager) initSQL(cfg DBConfig) *Errors.AppError {
	var sqlConfig SQLConfig
	b, _ := json.Marshal(cfg)
	err := json.Unmarshal(b, &sqlConfig)
	if err != nil {
		fmt.Println(err)
		return Errors.AppError{}.DbConnectionError("Invalid SQL config")
	}
	var errr *Errors.AppError
	var sqlDB Database = &SqlDatabase{}
	fmt.Println("asdfads", sqlDB)
	log.Printf("ptr address: %p", sqlDB)
	errr = sqlDB.Connect(sqlConfig)
	if errr != nil {
		fmt.Println(errr.Message)
		return errr
	}
	log.Printf("ptr address: %p", sqlDB)
	fmt.Println("asdfad", sqlDB)
	dc.DBClientsMap[sqlConfig.NAME] = sqlDB
	return nil
}
