package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncCondRegular(t *testing.T) {
	var cond = sync.NewCond(&sync.Mutex{})
	var group sync.WaitGroup

	group.Add(1)

	go func() {
		defer group.Done()
		cond.L.Lock()
		fmt.Println("Waiting for signal...")
		cond.Wait()
		fmt.Println("Received signal, proceeding with the task.")
		cond.L.Unlock()
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Sending signal to the goroutine...")
	cond.Signal()

	group.Wait()
	fmt.Println("TestSyncCond completed")
}

func TestSyncCondBroadcast(t *testing.T) {
	var cond = sync.NewCond(&sync.Mutex{})
	var group sync.WaitGroup

	for i := 0; i < 3; i++ {
		group.Add(1)

		go func(id int) {
			defer group.Done()
			cond.L.Lock()
			fmt.Printf("Goroutine %d waiting for signal...\n", id)
			cond.Wait()
			fmt.Printf("Goroutine %d received signal, proceeding with the task.\n", id)
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Sending broadcast signal to all goroutines...")
	cond.Broadcast()

	group.Wait()
	fmt.Println("TestSyncCondBroadcast completed")
}

func TestSyncCondMultipleSignals(t *testing.T) {
	var cond = sync.NewCond(&sync.Mutex{})
	var group sync.WaitGroup

	for i := 0; i < 3; i++ {
		group.Add(1)

		go func(id int) {
			defer group.Done()
			cond.L.Lock()
			fmt.Printf("Goroutine %d waiting for signal...\n", id)
			cond.Wait()
			fmt.Printf("Goroutine %d received signal, proceeding with the task.\n", id)
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		fmt.Printf("Sending signal to goroutine %d...\n", i)
		cond.Signal()
		time.Sleep(1 * time.Second)
	}

	group.Wait()
	fmt.Println("TestSyncCondMultipleSignals completed")
}
