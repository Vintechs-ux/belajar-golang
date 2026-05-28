package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func HitungSirkuit(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	return total
}

func DisplayNumber(n int) {
	fmt.Println("Display", n)
}

func TestHitungSirkuit(t *testing.T) {
	go fmt.Println(HitungSirkuit(1000))
	fmt.Println("Ups")
	time.Sleep(100 * time.Millisecond)
}

func TestHitungSirkuitSync(t *testing.T) {
	fmt.Println(HitungSirkuit(1000))
	fmt.Println("Ups")
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(2 * time.Second)
}
