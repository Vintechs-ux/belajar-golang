package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestDataRace(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var idr_balance int64 = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mutex.Lock()
				idr_balance += 1
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(idr_balance)
}
