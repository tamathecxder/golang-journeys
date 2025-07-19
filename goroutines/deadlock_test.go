package goroutines

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name   string
	Amount int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) UpdateAmount(amount int) {
	u.Amount = u.Amount + amount
}

func Transfer(fromUser *UserBalance, toUser *UserBalance, amount int) {
	fromUser.Lock()
	fromUser.UpdateAmount(-amount)
	fmt.Println("Locking User 1...")

	time.Sleep(1 * time.Second)

	toUser.Lock()
	toUser.UpdateAmount(amount)
	fmt.Println("Locking User 2...")

	time.Sleep(1 * time.Second)

	fromUser.Unlock()
	toUser.Unlock()
}

func TestDeadlock(t *testing.T) {
	asep := UserBalance{
		Name:   "Asep",
		Amount: 5000,
	}

	ujang := UserBalance{
		Name:   "Ujang",
		Amount: 5000,
	}

	go Transfer(&asep, &ujang, 1000)
	go Transfer(&ujang, &asep, 2000)

	time.Sleep(5 * time.Second)

	fmt.Println("Transfer process is completed")
	fmt.Printf("%s total balance: %s\n", asep.Name, strconv.Itoa(asep.Amount))
	fmt.Printf("%s total balance: %s\n", ujang.Name, strconv.Itoa(ujang.Amount))
}

func AsyncCall(group *sync.WaitGroup, result string) {
	defer group.Done()

	group.Add(1)

	fmt.Println(result)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go AsyncCall(group, "This is a counter on process: "+strconv.Itoa(i))
	}

	group.Wait()

	fmt.Println("The whole process is completed!")
}

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println("Total counter:", counter)
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "DEFAULT_VALUE"
		},
	}
	group := sync.WaitGroup{}

	firstStr := "Hello"
	secondStr := "World"

	pool.Put(&firstStr)
	pool.Put(&secondStr)

	for range 10 {
		group.Add(1)

		go func() {
			res := pool.Get()

			if str, ok := res.(*string); ok {
				fmt.Println("response:", *str)
			}

			time.Sleep(10 * time.Second)

			pool.Put(res)
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println("Pool is completed")
}

func TestSyncMap(t *testing.T) {
	var data sync.Map
	var group sync.WaitGroup

	for i := 1; i <= 100; i++ {
		group.Add(1)

		go func(i int) {
			defer group.Done()
			data.Store("key-"+strconv.Itoa(i), "value-"+strconv.Itoa(i))
			fmt.Println("Data stored for key:", "key-"+strconv.Itoa(i))
		}(i)
	}

	group.Wait()
	fmt.Println("SyncMap process is completed")
	data.Range(func(key, value any) bool {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
		return true
	})

	fmt.Println("All data in SyncMap has been printed")
	fmt.Println("TestSyncMap completed")
}

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
