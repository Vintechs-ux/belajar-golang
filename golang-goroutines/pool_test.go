package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	var wg sync.WaitGroup

	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}

	pool.Put("Diki")
	pool.Put("Surya")
	pool.Put("Saputra")

	data1 := pool.Get()
	fmt.Println("diluar go: ", data1)
	data2 := pool.Get()
	fmt.Println("diluar go: ", data2)
	pool.Put(data1)
	pool.Put(data2)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	wg.Wait()
	fmt.Println("Proses selesai")
}
