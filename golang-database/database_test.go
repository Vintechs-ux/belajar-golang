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

	_, err := db.ExecContext(ctx, script, "laura")
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses insert customer baru")
}

func TestQueryPostgres(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("=== ISI TABEL CUSTOMER ===")

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d | Name: %s\n", id, name)

	}
}

type Customer struct {
	id   int
	name string
}

func ForBool() bool {
	fmt.Println("Inisialisasi")
	sliceCustomer := make([]Customer, 0, 10)
	Siregar := Customer{1, "Siregar"}
	Budi := Customer{2, "Budi"}
	Kosong := Customer{0, ""}

	sliceCustomer = append(sliceCustomer, Siregar)
	sliceCustomer = append(sliceCustomer, Budi)
	sliceCustomer = append(sliceCustomer, Kosong)
	slice := []struct{ Customer }{
		{
			Siregar,
		},
		{
			Budi,
		},
		{
			Kosong,
		},
	}

	fmt.Println(sliceCustomer)
	fmt.Println(len(sliceCustomer))
	fmt.Println(cap(sliceCustomer))
	fmt.Println(slice)

	return true
}

func TestForBool(t *testing.T) {
	ForBool()
}
