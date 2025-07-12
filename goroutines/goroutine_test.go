package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func CurrentTime() {
	fmt.Println("It's 10 PM")
}

func TestGoroutine(t *testing.T) {
	go CurrentTime()
	fmt.Println("Print from TestGoroutine")

	time.Sleep(3 * time.Second)
}
