package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Complete")
}

var counting = 0

func OnlyOnce() {
	counting++
}

func TestOnlyOnce(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			once.Do(OnlyOnce)
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counting)
}

var counter = 0

func Counter() {
	counter++
}

func TestRaceWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			Counter()
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func TestRaceWaitGroupLoop(t *testing.T) {
	for putaran := 1; putaran <= 1000; putaran++ {
		t.Run(fmt.Sprintf("Putaran-%d", putaran), func(t *testing.T) {
			counter := 0
			var wg sync.WaitGroup
			var mutex sync.Mutex

			for i := 0; i < 100; i++ {
				mutex.Lock()
				wg.Add(1)
				go func() {
					defer wg.Done()
					counter++
					mutex.Unlock()
				}()
			}
			wg.Wait()

			assert.Equal(t, 100, counter, "program seharusnya tidak mengalami race condition")
		})
	}
}
