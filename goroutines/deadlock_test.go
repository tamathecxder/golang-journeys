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
