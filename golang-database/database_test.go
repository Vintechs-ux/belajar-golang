package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestConnection(t *testing.T) {
	connStr := "user=postgres password=rahasia dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal mendapat koneksi ke database", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Tidak ada respon dari database", err)
	}

	fmt.Println("Koneksi Postgresql sukses!")

	query := `CREATE TABLE IF NOT EXISTS customer (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL

	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Gagal membuat tabel customer", err)
	}
	fmt.Println("Tabel 'customer' berhasil dibuat")
}

func TestGetConnection(t *testing.T) {
	GetConnection()
}

func TestExecPostgres(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(name) VALUES($1)"

	_, err := db.ExecContext(ctx, script, "budi")
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses insert customer baru")
}
