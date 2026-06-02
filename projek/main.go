package main

import (
	"fmt"
	"sync"
	"time"
)

type GasLog struct {
	DeviceID string
	GasLevel int
	Timestamp time.Time
}

func main() {

	gasLogPool := &sync.Pool{
		New: func ()  any {	
			return &GasLog{}
		}
	}

	var wg sync.WaitGroup
	devices := []string{"ESP32_ROOM_A", "ESP32_ROOM_B", "ESP32_ROOM_C"}
	fmt.Println("=== MEMULAI SIMULASI MONITORING GAS GOROUTINES")

	for i := 1; i <= 5; i++{
		wg.Add(1)
		go func (workerID int)  {
			logData := gasLogPool.Get().(*GasLog)

		}()
	}


}
