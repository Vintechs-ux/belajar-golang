package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Heru"
		fmt.Println("Selesai mengirim data")
	}()

	func() {
		fmt.Println("Ini anonymus function")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Heru Kurniawan"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Heru Kurniawan"
	fmt.Println("Raditya")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	defer close(channel)

	channel <- "Data emas 1"
	channel <- "Data emas 2"
	channel <- "Data emas 3"

	fmt.Println("Kapasitas Maksimal Buffer: ", cap(channel))
	fmt.Println("Jumlah Data Buffer Saat ini:  ", len(channel))

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Sisa Data di Buffer Setelah Diambil: ", len(channel))
}

func TestRangeChannel(t *testing.T) {
	queue := make(chan string, 5)

	go func() {
		for i := 0; i <= 5; i++ {
			queue <- fmt.Sprintf("Log transaksi ke %d", i)
		}

		close(queue)
		fmt.Println("=== Sirkuit Pengirim Selesai & Channel  Ditutup ===")
	}()

	for data := range queue {
		fmt.Println("Penerima berhasil menyedot", data)
	}

	fmt.Println("=== Semua Data Berhasil Diproses Tanpa Deadlock ===")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2: ", data)
			counter++
		default:
			fmt.Print("\rMemeriksa transaksi, mohon tunggu")
		}
		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	saldo   int
}

func (account *BankAccount) UbahSaldo(jumlah int) {
	account.RWMutex.Lock()
	account.saldo += jumlah
	account.RWMutex.Unlock()
}

func (account *BankAccount) LihatSaldo() int {
	account.RWMutex.RLock()
	value := account.saldo
	account.RWMutex.RUnlock()
	return value
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.UbahSaldo(1)
				fmt.Println(account.LihatSaldo())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance : ", account.LihatSaldo())
}

type SmartLogger struct{}

func (l *SmartLogger) GetLog() bool {
	fmt.Println("Sirkuit Logika aman , menginisialisasikan program selanjutnya...")
	return true
}

func TestSmartLogger(t *testing.T) {
	var logger SmartLogger
	result := logger.GetLog()
	assert.Equal(t, true, result, "Logika Logger Error")
}

type UserBalance struct {
	sync.Mutex
	name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.name)
	user1.Change(amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlockTransfer(t *testing.T) {
	budi := UserBalance{name: "Budi", Balance: 100}
	diki := UserBalance{name: "Diki", Balance: 200}

	fmt.Println("Budi: ", budi)
	fmt.Println("Diki: ", diki)

	go Transfer(&budi, &diki, 500)
	go Transfer(&diki, &budi, 600)

	time.Sleep(5 * time.Second)

	fmt.Println("Budi: ", budi)
	fmt.Println("Diki: ", diki)
	fmt.Println("Hai")
}
