package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// Init ...
func Init() {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"))
	db, err = sql.Open("mysql", connectString)
	if err != nil {
		// Error.Error(err)
		panic("DB Connection Error")
	}

	err = db.Ping()
	if err != nil {
		// Error.Error(err)
		panic("Ping DB Connection Error")
	}

	// db.SetMaxOpenConns()                // Max connection
	// db.SetMaxIdleConns()                 // Max idle connection
	// db.SetConnMaxLifetime(3 * time.Second) // Max lifetime
}

// GetDB ...
func GetDB() *sql.DB {
	return db
}
