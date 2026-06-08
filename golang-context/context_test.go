package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextA, "g", "G")

	slice := []context.Context{contextA, contextB, contextC, contextD, contextE, contextF, contextG}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println(contextF.Value("c"))

	fmt.Println(contextG.Value("g"))
	fmt.Println(contextE.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func CreateCounter(ctx context.Context) (chan int, *sync.WaitGroup) {
	destination := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++

			}
		}
	}()
	return destination, wg
}

func TestCounter(t *testing.T) {
	fmt.Println("Total goroutines: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination, wg := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n%100 == 0 {
			break
		}
	}
	cancel()
	wg.Wait()

	fmt.Println("Total goroutines now: ", runtime.NumGoroutine())
}
