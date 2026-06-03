package main

import (
	"fmt"
	"sync"
)

func main() {
	pingChan := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("[Goroutines] Mengirim data ke pipa...")
		pingChan <- "PING! Paket dari seberang"
		fmt.Println("[Goroutines] Data sukses dijemput, program berlanjut!")
	}()

	fmt.Println("[Main] Menunggu jemputan data")
	data := <-pingChan

	wg.Wait()

	fmt.Printf("[Main] Data diterima: %s\n", data)

	gudangChan := make(chan int, 2)

	fmt.Println("Data:", len(gudangChan))
	fmt.Println("Kapasitas:", cap(gudangChan))

	gudangChan <- 10
	gudangChan <- 11

	fmt.Println("Data:", len(gudangChan))
	fmt.Println("Kapasitas:", cap(gudangChan))

	fmt.Println(<-gudangChan)
	fmt.Println(len(gudangChan))

	sharedChan := make(chan int, 3)

	go Produser(sharedChan)
	Consumer(sharedChan)
}

func Produser(ch chan<- int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}

func Consumer(ch <-chan int) {
	for data := range ch {
		fmt.Printf("[Consumer] Mendapat Angka: %d\n", data)
	}
}
