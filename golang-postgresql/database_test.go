package main

import (
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestOpenConnection(t *testing.T) {
	GetConnection()
}

func TestMain(t *testing.T) {
	fmt.Println("hai")
}
