package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)

	fmt.Println(time.Now())

	res := <-timer.C

	fmt.Println(res)
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(3 * time.Second)

	fmt.Println(time.Now())

	res := <-channel

	fmt.Println(res)
}

func TestTimerFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Mikum")

		group.Done()
	})

	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(2 * time.Second)

		ticker.Stop()
	}()

	for v := range ticker.C {
		fmt.Println("Value:", v)
	}

	fmt.Println("Ticker is stopped")
}

func TestTickOnly(t *testing.T) {
	tick := time.Tick(1 * time.Second)

	go func() {
		time.Sleep(2 * time.Second)
	}()

	for v := range tick {
		fmt.Println("Value:", v)
	}

	fmt.Println("Ticker is stopped")
}
