package databasemysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_golang_journeys")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
