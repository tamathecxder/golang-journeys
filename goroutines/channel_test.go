package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	strChan := make(chan string)

	defer close(strChan)

	go func() {
		time.Sleep(2 * time.Second)
		strChan <- "EXAMPLE"
	}()

	data := <-strChan

	fmt.Println("Data from channel:", data)
}
