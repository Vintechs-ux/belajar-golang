package golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCpu)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads: ", totalThreads)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines: ", totalGoroutines)

	wg.Wait()
}
