package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456789@tcp(localhost:3306)/golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate -database "mysql://root:123456789@tcp(localhost:3306)/golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:123456789@tcp(localhost:3306)/golang_database_migration" -path db/migrations down
}
