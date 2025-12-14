package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	Errors "github.com/nurulnabi/go-finsight/internal/errors"
)

type SqlDatabase struct {
	//SqlDatabase this an struct representing one instance of an sql connection
	name string
	db   *sql.DB
}

// Connect establishes a connection to the database, and verifies if the database is connected successfuly
//
// Parameters:
//   - cfg: database configuration.
//
// Returns:
//   - *Errors.AppError: non-nil on failure
func (sdb *SqlDatabase) Connect(dbConfig DBConfig) error {
	fmt.Println("SqlDatabase.Connect start")
	sqlConfig, ok := dbConfig.(SQLConfig)
	if !ok {
		err := Errors.AppError{}.DbConnectionError("SqlDatabase.Connect SQL Config parsing failed")
		fmt.Println(err)
		return err
	}
	if sqlConfig.DB_URI != "" {
		db, err := sql.Open("postgres", sqlConfig.DB_URI)
		if err != nil {
			return Errors.AppError{}.DbConnectionError(err.Error())
		}
		fmt.Printf("SqlDatabase.Connect conn creation sucessful")

		if err := db.Ping(); err != nil {
			return Errors.AppError{}.DbConnectionError(err.Error())
		}
		fmt.Printf("SqlDatabase.Connect conn verified")

		sdb.db = db
		sdb.name = sqlConfig.NAME
	} else {
		fmt.Printf("SqlDatabase.Connect sqlConfig.DB_URI not found %v", sqlConfig.DB_URI)
	}
	return nil
}

// HealthCheck checks if the database is still connected
//
// Parameters:
//   - ctx: controls cancellation and timeout.
//
// Returns:
//   - *Errors.AppError: non-nil on failure
func (sdb *SqlDatabase) HealthCheck(ctx context.Context) error {
	if err := sdb.db.PingContext(ctx); err != nil {
		return Errors.AppError{}.DbConnectionError(err.Error())
	}
	fmt.Printf("SqlDatabase.HealthCheck sucessful")
	return nil
}

func (sdb *SqlDatabase) Close() error {
	if err := sdb.db.Close(); err != nil {
		fmt.Printf("SqlDatabase.Close failed %v", err)
		return Errors.AppError{}.DbConnectionError(err.Error())
	}
	fmt.Printf("SqlDatabase.Close conn closed successful")
	return nil
}
