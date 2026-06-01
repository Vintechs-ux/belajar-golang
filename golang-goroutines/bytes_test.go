package golang_goroutines

import (
	"fmt"
	"testing"
)

func TestLihatBytes(t *testing.T) {
	data := "Halo"
	bytes := []byte(data)

	fmt.Println(bytes)
	fmt.Println(bytes[0])
	fmt.Println(string(bytes))
}
