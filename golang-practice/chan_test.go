package main

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 5; i++ {
			channel <- i
		}
	}()
	for data := range channel {
		fmt.Printf("Data: %d\n", data)
	}
}

func TestChan2(t *testing.T) {
	ch := make(chan int, 3)

	for i := 0; i <= 10; i++ {
		select {
		case ch <- i:
			fmt.Printf("Sent: %d\n", i)
		default:
			fmt.Printf("Skipped: %d\n", i)

		}
	}
}

func TestChan3(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2

	for i := 0; i <= 5; i++ {
		select {
		case data := <-ch:
			fmt.Printf("Sent: %d\n", data)
		default:
			fmt.Printf("empy skip\n")
		}
	}
}
