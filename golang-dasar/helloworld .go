package main

import (
	"fmt"
)

var (
	a      uint8 = 25
	b      uint8 = 25
	angka1       = 3
	angka2       = 5
)

const (
	firstName = "Budi"
	lastName  = "Nugraha"
)

var (
	total    uint8  = a + b
	message  string = "total lebih besar dari 30"
	message2 string = "total lebih kecil dari 30"
)

type (
	NoKTP     string
	Filter    func(string) string
	BlackList func(string) bool
)

////// Function

func sayHello() {
	fmt.Println("Halo")
}

func Hitung(Angka1 int, Angka2 int) (int, int) {
	return Angka1 + Angka2, Angka1 * Angka2
}

func getFullName() (firstname, middlename, lastname string) {
	firstname = "Eko"
	middlename = "Kurniawan"
	lastname = "Kanedi"

	return firstname, middlename, lastname
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func sumAllSliceParam(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func getGoodBye(name string) string {
	return "Selamat Tinggal " + name
}

var FilterNama Filter = func(nama string) string {
	if nama == "Anjing" {
		return "..."
	} else {
		return nama
	}
}

var FilterPassword Filter = func(string) string {
	return "*********"
}

func sayHi() {
	fmt.Println("Hai")
}

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func inputPasswordWithFilter(password string, filter Filter) {
	fmt.Println("Password: ", filter(password))
	fmt.Println("Save to database")
}

func sensor(password string) string {
	return "********"
}

func registerUser(name string, blacklist BlackList) {
	defer sayHi()
	if blacklist(name) {
		fmt.Println("You Are Blocked!!!")
	} else {
		fmt.Println("User Registered")
	}
}

func checkBlocked(name string) bool {
	if name == "Anjing" {
		return true
	} else {
		return false
	}
}

func factorialLoop(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialRecursive(number-1)
	}
}

func endApp() {
	fmt.Println("Program berhenti Error")
	pesan := recover()
	fmt.Println("Terjadi Error: ", pesan)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

///// STRUCT

type Customer struct {
	Name, Address string
	Age           int
}

type Person struct {
	Name string
}

type HasName interface {
	getName() string
}

type Address struct {
	City     string
	Country  string
	Province string
}

type ID struct {
	Nama  string
	Kelas int
	Umur  int
}

type Chat struct {
	ID   int
	Nama string
	Time int
}

func (customer Customer) sayHai() {
	fmt.Println("HALOOO", customer.Name)
}

func (person Person) getName() string {
	return person.Name
}

func siapaKamu(value HasName) {
	fmt.Println("Halo", value.getName())
}

////// Pointer 2 /////////////////////////////////////////////////

func changeTime(chat *Chat) {
	chat.Time = 12
}

type Man struct {
	Name string
}

func (man *Man) Married(cincin bool) {
	man.Name = "Mr " + man.Name
	m := man.Name
	if m == man.Name || cincin {
		fmt.Println("Lah sudah nikah bang " + m)
	} else {
		fmt.Println("Gak mungkin belum nikah kocak")
	}
}

///// Interface Kosong

func Ups() any {
	return "ups"
}

///// Main Function

func main() {
	fmt.Println("Hello World")

	var umur int = 18

	fmt.Println(umur)

	umur = 20

	fmt.Println(umur)

	if total > 30 {
		fmt.Println(message, len(message), firstName, lastName, NoKTP("123456789"))
	} else {
		fmt.Println(total, message2, len(message2), firstName, lastName)
	}

	angka1 += 10

	fmt.Println(angka1)

	angka2++

	fmt.Println(angka2)

	if angka1 > 1 && angka2 > 5 {
		fmt.Println(angka1, "lebih besar dari 1 dan", angka2, "lebih besar dari 5")
	}

	arr := [...]int{1, 2, 2, 3, 4}

	i := 2

	fmt.Println(arr[i])

	for index, value := range arr {
		fmt.Println(index, value)
	}

	name := [...]string{"siregar", "budiman", "manalu", "putra", "sitompul", "sijeki"}

	slice := name[:2]

	fmt.Println(slice)

	hari := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice := hari[5:]
	fmt.Println(daySlice)
	daySlice[0] = "senin baru"
	daySlice[1] = "selasa baru"

	fmt.Println(daySlice[0], daySlice[1])
	fmt.Println(daySlice)

	daySlice2 := append(daySlice, "liburr")
	fmt.Println(daySlice2)

	daySlice2[0] = "Ups"

	fmt.Println(daySlice)
	fmt.Println(daySlice2)
	fmt.Println(hari)

	fromSlice := hari[:]

	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	person := map[string]string{
		"nama": "raditya",
		"umur": "19",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["umur"])

	length := len(person)

	fmt.Println(length)

	switch length {
	case 2:
		fmt.Println("valid")
	case 3:
		fmt.Println("siregar")
	}

	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke:", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 20; counter2++ {
		fmt.Println("Perulangan ke:", counter2)
	}

	namess := []string{"eko", "kurniawan", "sitompul"}

	for i := 0; i < len(namess); i++ {
		fmt.Println(namess[i])
	}

	for _, name := range namess {
		fmt.Println(name)
	}

	status_Loading := false
	counter3 := 1
	Maks := 20

	for counter3 <= Maks {
		if counter3 == 10 && status_Loading {
			fmt.Println("BATAS MAKSIMUM ")
			break
		} else {
			fmt.Println("Perulangan ke:", counter3)
			counter3++
		}
	}

	for i := 0; i < 20; i++ {
		jule := 1
		if i%2 == 0 {
			jule += i
			continue
		}

		pertambahan, perkalian := Hitung(i, jule)

		fmt.Println("perulangan ke (ganjil):", i)
		sayHello()
		fmt.Println(pertambahan)

		if perkalian <= 20 {
			firstName, middleName, lastName := getFullName()
			fmt.Println(firstName, middleName, lastName)
			total := sumAll(i, jule, i, jule, i, jule)
			fmt.Println(total)

			numberSlice := []int{i, i, i, jule, jule, jule}
			totalSlice := sumAllSliceParam(numberSlice...)
			fmt.Println(totalSlice)
		}

		goodbye := getGoodBye
		fmt.Println(goodbye("Siregar"))
		sayHelloWithFilter("Anjing", FilterNama)
		inputPasswordWithFilter("haloereijiog", FilterPassword)
		checkBlock := checkBlocked

		registerUser("Anjing", checkBlock)

	}
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(10))

	runApp(true)
	fmt.Println("Program Berlanjut")

	var eko Customer
	eko.Name = "eko"
	eko.Address = "Surabaya"
	eko.Age = 19
	fmt.Println(eko)

	siregar := Customer{
		Name:    "siregar",
		Address: "cuiarta",
		Age:     99,
	}

	fmt.Println(siregar)

	owo := Customer{"dodol", "didil", 12}
	fmt.Println(owo)

	rully := Customer{"rully", "tulindo", 14}
	rully.sayHai()

	joko := Person{"raditya"}
	siapaKamu(joko)
	fmt.Println(Ups())

	Alamat := Address{"Jakarta", "Indonesia", "Jawa Barat"}
	fmt.Println(Alamat)
	Alamat2 := Alamat
	Alamat2.City = "Miekarta"
	fmt.Println(Alamat)
	fmt.Println(Alamat2)

	// POINTER

	raditya := ID{"Raditya", 12, 19}
	fmt.Println(raditya)
	raditya2 := &raditya
	raditya2.Umur = 20

	fmt.Println(raditya2)

	// raditya2 := &ID{"Adit", 13, 25}
	*raditya2 = ID{"Ditompul", 12, 23}

	fmt.Println(raditya)
	fmt.Println(raditya2)

	raditya2.Nama = "siregar"
	fmt.Println(raditya)
	fmt.Println(raditya2)

	//////// POINTER 2

	budi := Chat{1, "budi", 13}
	fmt.Println(budi)
	changeTime(&budi)
	fmt.Println(budi)

	judi := Man{"Judi"}
	judi.Married(false)
	fmt.Println(judi)
}
