package app

import (
	"database/sql"
	"fmt"
	"golangeko/helper"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/golang_eko?sslmode=disable")

	helper.PanicIfError(err)

	fmt.Println("Koneksi ke database berhasil")
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
