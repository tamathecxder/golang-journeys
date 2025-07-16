package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var mutex sync.Mutex
	counter := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				counter = counter + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final Result:", counter)
}

type BankAccount struct {
	Mutex   sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.Mutex.Lock()
	account.Balance = account.Balance + amount
	account.Mutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.Mutex.RLock()
	balance := account.Balance
	account.Mutex.RUnlock()

	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for range 100 {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Total balance", account.GetBalance())
}
