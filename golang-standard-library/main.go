package main

import (
	"bufio"
	"container/list"
	"container/ring"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("notfound error")
)

func GetByID(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "Raditya" {
		return NotFoundError
	}

	return nil
}

type VolumeControl struct {
	level int
}

type Health func(int) (string, int)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var CheckHealth Health = func(index int) (string, int) {
	if index <= 50 {
		return "Sound Tidak Sehat, Mengurangi Volume Sebanyak Setengah", 0
	} else if index <= 60 {
		return "Sound Dibawah Batas Normal, Mengurangi Volume Sebanyak 40 Persen", 2
	} else {
		return "Sound Sehat", 1
	}
}

func (v *VolumeControl) SetLevel(value int, index int, health Health) {
	str, check := health(index)
	fmt.Println(check)
	if check == 0 {
		fmt.Println(str)
		v.level = max(value*1/2, 0)
	}

	if check == 2 {
		fmt.Println(str)
		v.level = max(value*6/10, 0)
	}

	if check == 1 {
		fmt.Println(str)
		v.level = max(value, 0)
	}

	if v.level > 100 {
		v.level = 100
	}
}

func (v *VolumeControl) Mute(value int) {
	v.level = 0
}

type Sample struct {
	Name    string `required:"true" max:"10"`
	Age     int    `required:"true" max:"2"`
	Address string `required:"false" max:"100"`
}

func RedField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println(valueType)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type, "required:", structField.Tag.Get("required"))
	}
}

type Orang struct {
	Name string
}

func (o *Orang) ChangeName(name string) {
	o.Name = "Siregar"
}

type Users struct {
	Name    string
	Age     int
	Address string
}

func (u *Users) getName() (index string, jarak string) {
	if u.Age < 10 && u.Address == "Surabaya" {
		index, jarak = "Tidak Masuk Kategori", "Jauh"
		return
	} else {
		return "Hai", "kamu di dis"
	}
}

func SaveToDatabase(name *string) {
	*name = "Data " + *name + " Berhasil di simpan ke DB!"
}

func CekError(name string) (status string, err error) {
	if name == "" {
		status, err = "Gagal", errors.New("nama tidak boleh kosong")
		return
	}
	status, err = "Berhasil", nil
	return
}

type Uwong struct {
	Name    string
	Age     int
	Address string
}

func (u *Uwong) GetName() (index int, err error) {
	if u.Name == "anjing" {
		u.Name = "******"
		index, err = 400, errors.New("nama tidak pantas")
		return
	}
	index, err = 200, nil
	return
}

func TesError() (str string, int int) {
	str, int = "Berubah", 1
	return
}

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line)

	}
	return message, err
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(message)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Hello World")

	firstname := "siregar"
	lastname := "sitompul"

	fmt.Printf("Hello %s %s\n", firstname, lastname)

	err := GetByID("Raditya")

	cobaCek := errors.Is(err, ValidationError)
	fmt.Println(cobaCek)
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		}
	} else {
		fmt.Println("Sukses")
	}

	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println(err)
	}

	sound := VolumeControl{level: 0}
	sound.SetLevel(33, 30, CheckHealth)
	fmt.Println(sound.level)

	name := flag.String("nama", "World", "Nama orang")

	flag.Parse()

	fmt.Printf("Halo %s , senang bertemu dengan mu\n", *name)

	msg := "Halo sayang"
	fmt.Println(msg)
	msg = strings.ToLower(msg)
	fmt.Println(msg)
	msg = strings.Trim(msg, "sayang")
	fmt.Println(msg)

	i, err := strconv.Atoi("2423")
	if err == nil {
		fmt.Println(i)
	} else {
		fmt.Println("error", err)
	}

	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Kaneddy")

	fmt.Println(data.Front())
	fmt.Println(data)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	ContohSlice := []string{"Raditya", "Siregar", "Sitompul"}
	fmt.Println(ContohSlice)
	ContohSlice2 := make([]string, 3, 5)
	fmt.Println(len(ContohSlice2), cap(ContohSlice2))
	ContohSlice2 = append(ContohSlice2, "Sitompul")

	ContohSlice = append(ContohSlice, ContohSlice2...)
	fmt.Println(ContohSlice)

	//// Closure

	buatAntrian := func() func() int {
		nomor := 0
		return func() int {
			nomor++
			return nomor
		}
	}

	antreanPasien := buatAntrian()

	fmt.Println(antreanPasien())
	fmt.Println(antreanPasien())
	fmt.Println(antreanPasien())

	dataRing := ring.New(5)
	fmt.Println(dataRing.Len())

	for i := 0; i < dataRing.Len(); i++ {
		dataRing.Value = "Value " + strconv.FormatInt(int64(i), 10)
		dataRing = dataRing.Next()
	}

	dataRing.Do(func(value any) {
		fmt.Println(value)
	})

	users := []User{
		{"Eko", 30},
		{"Budi", 25},
		{"Siregar", 43},
		{"Joko", 4343},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

	for i, now := 0, time.Now(); i < 10; i++ {
		fmt.Println(now.Local())
	}

	duration := 100 * time.Second
	fmt.Println(duration)
	fmt.Printf("%s\n", duration)

	RedField(Sample{"siregar", 21, "Surabaya"})

	/////////////// REGEXP /////////////////////////////////////

	regex := regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eKo"))
	fmt.Println(regex.MatchString("ewo"))
	fmt.Println(regex.MatchString("siregar"))

	fmt.Println(regex.FindAllString("eko edi siregar elo eso elo ego eno bening eto", 10))

	/////////////// ENCODING ////////////////////////////////////

	nome := "Siregar Setiawan Budi Sitompul"

	encoded := base64.StdEncoding.EncodeToString([]byte(nome))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	csvString := "eko,kurniawan,khanedy\n" + "budi,pratama,sitompul\n" + "siregar,sitompul,sijule\n"
	reader := csv.NewReader(strings.NewReader(csvString))
	fmt.Println(reader)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	ditompul := Users{"ditompul", 2, "Surabaya"}
	index, jarak := ditompul.getName()
	fmt.Println(index, jarak)

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Siregar", "sitompul", "sibujang"})
	_ = writer.Write([]string{"Sijule", "sitompul", "sibujang"})
	_ = writer.Write([]string{"Sibugar", "sitompul", "sibujang"})
	_ = writer.Write([]string{"Simalu", "sitompul", "sibujang"})
	writer.Flush()

	nomes := []string{"Rully", "Rafa", "Dika", "Kanedi", "Surya"}
	nomburs := []int{2, 5, 2, 65}
	contain := slices.Contains(nomes, "Surya")
	fmt.Println(contain)
	slices.Sort(nomburs)
	fmt.Println(nomburs)
	fmt.Println(slices.Max(nomburs))

	Ditoon := Orang{"Ditoon"}
	fmt.Println(Ditoon)
	Ditoon.ChangeName("Siregar")
	fmt.Println(Ditoon.Name)

	userBaru := "Raditya"
	SaveToDatabase(&userBaru)
	fmt.Println(userBaru)
	fmt.Println("Halo dunia")
	nomus := "ditok"
	sapa := "Hai" + nomus
	fmt.Println(sapa)
	fmt.Println("Hai nama saya raditya", nomus)

	fmt.Println(path.Dir("/home/vintechs/Projects/GO/belajar-golang-standard-library/main.go"))
	fmt.Println(path.Base("/home/vintechs/"))
	fmt.Println(path.Base("/home/vintechs/Documents/Golang/Golang-Dasar/Salinan dari Go-Lang Dasar.pptx"))
	fmt.Println(path.Ext("/home/vintechs/Games/darksouls/dosdevices/68ec4fa4-16a3-4ba9-939e-3abf5d97b061"))
	fmt.Println(path.Join("home", "vintechs", "Games"))

	fmt.Println(filepath.Dir("/home/vintechs/Projects/GO/belajar-golang-standard-library/main.go"))
	fmt.Println(filepath.Base("/home/vintechs/Projects/GO/belajar-golang-standard-library/main.go"))
	fmt.Println(filepath.Ext("/home/vintechs/Projects/GO/belajar-golang-standard-library/main.go"))
	fmt.Println(filepath.Join("home", "vintechs", "Projects", "GO", "belajar-golang-standard-library", "main.go"))

	folderTarget := "."
	fmt.Println("Memulai scan file")

	err = filepath.WalkDir(folderTarget, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Printf("[Folder] %s\n", path)
		} else {
			if filepath.Ext(path) == ".go" {
				fmt.Printf("[File Go] %s -> Nama File: %s\n", path, d.Name())
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("Waduh Scan gagal: ", err)
	}

	statuscheck, err := CekError("Hai")
	fmt.Println(statuscheck, err)

	jeky := Uwong{"Jeky", 14, "Surabaya"}
	fmt.Println(jeky)
	fmt.Println(jeky.GetName())

	iniString := "Hai"
	fmt.Println(iniString)

	iniString, iniInt := TesError()
	fmt.Println(iniString, "Status: ", iniInt)

	input := strings.NewReader("halo nama saya raditya.\nSaya seorang kernel engineer\n")
	readers := bufio.NewReader(input)
	for {
		line, _, err := readers.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	writers := bufio.NewWriter(os.Stdout)
	_, _ = writers.WriteString("Hello World\n")
	_, _ = writers.WriteString("Hello GO\n")
	writers.Flush()
	// createNewFile("siregar.txt", "siregar")

	//result, _ := readFile("main.go")
	//fmt.Println(result)
	fmt.Println(addToFile("siregar.txt", "sitompul"))
}
