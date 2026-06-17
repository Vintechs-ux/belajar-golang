package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

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

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql := "SELECT id, name, email , balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		var balance float64
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Printf("ID       : %d\n", id)
		fmt.Printf("Name     : %s\n", name)
		fmt.Printf("Email    : %s\n", email)
		fmt.Printf("Balance  : %f\n", balance)
		fmt.Printf("Rating   : %.1f\n", rating)
		fmt.Printf("BirthDate: %s\n", birthDate.Format("2006-01-02"))
		fmt.Printf("Married  : %v\n", married)
		fmt.Printf("CreatedAt: %s\n", createdAt.Format("2006-01-02 15:04:05"))
		fmt.Println("----------------------------")
	}
}

func TestNullableSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	post := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, post)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email sql.NullString
		var balance float64
		var rating float64
		var birthDate sql.NullTime
		var married sql.NullBool
		var createdAt time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID       : %d\n", id)
		fmt.Printf("Name     : %s\n", name)
		if email.Valid {
			fmt.Printf("Email    : %s\n", email.String)
		}
		fmt.Printf("Balance  : %f\n", balance)
		fmt.Printf("Rating   : %.1f\n", rating)

		if birthDate.Valid {
			fmt.Println("BirthDate: ", birthDate.Time)
		}

		if married.Valid {
			fmt.Printf("Married  : %v\n", married.Bool)
		}

		fmt.Printf("CreatedAt: %s\n", createdAt.Format("2006-01-02 15:04:05"))
		fmt.Println("----------------------------")

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

func TestHelloWorld(t *testing.T) {
	fmt.Println("Hello World")
}
