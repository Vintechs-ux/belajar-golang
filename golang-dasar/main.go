package main

import (
	"belajar-golang-dasar/database"
	"belajar-golang-dasar/helper"
	_ "belajar-golang-dasar/internal"
	"errors"
	"fmt"
)

type notFoundError struct {
	Message string
}

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

func (e *notFoundError) Error() string {
	return e.Message
}

func Pembagian(value1, value2 int) (int, error) {
	if value2 == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return value1 / value2, nil
	}
}

func saveData(name string, data any) error {
	if name == "" {
		return &validationError{Message: "Validation error"}
	}
	if name != "eko" {
		return &notFoundError{Message: "Notfound error"}
	}

	return nil
}

func main() {
	result := helper.SayHello("Joko")
	fmt.Println(result)
	fmt.Println(helper.Version)
	fmt.Println(database.GetDatabase())

	fmt.Println("PEMBAGIAN")
	fmt.Println(Pembagian(4, 0))
	fmt.Println("PEMBAGIAN KE DUA")
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println(saveData("", 100))

}
