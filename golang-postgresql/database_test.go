package main

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestOpenConnection(t *testing.T) {
	dsn := "postgres://postgres:rahasia@localhost:5432/belajar_golang?sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
