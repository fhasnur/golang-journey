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

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third

	// Migration to a spesific version
	// migrate -database "mysql://root:123456789@tcp(localhost:3306)/golang_database_migration" -path db/migrations up [number] (example: 1, 2, 3,..)
	// migrate -database "mysql://root:123456789@tcp(localhost:3306)/golang_database_migration" -path db/migrations down [number] (example: 1, 2, 3,..)
}
