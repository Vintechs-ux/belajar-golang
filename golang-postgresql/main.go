package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "postgres://postgres:rahasia@localhost:5432/belajar_golang?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
}
