package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCondSimple(t *testing.T) {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	cond := sync.NewCond(&mutex)
	dataReady := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		mutex.Lock()
		if !dataReady {
			fmt.Println("Data belum ada, proses ditidurkan...")
			cond.Wait()
		}
		fmt.Println("Proses dibangunkan , menyiapkan data")
		mutex.Unlock()
	}()

	time.Sleep(2 * time.Second)
	mutex.Lock()
	dataReady = true
	fmt.Println("Producer: data siap!")
	cond.Signal()
	mutex.Unlock()
	wg.Wait()
	fmt.Println("Proses berjalan dengan baik")
}
