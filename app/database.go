package app

import (
	"database/sql"
	"golangeko/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("pgsql", "postgres:postgres@tcp(localhost:5432)/golang_eko")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
