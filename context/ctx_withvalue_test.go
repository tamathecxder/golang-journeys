package context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

type key string

func TestContextWithValue(t *testing.T) {
	var countryKey key = "country"
	var provinceKey key = "province"
	var districtKey key = "district"
	var subdistrictKey key = "subdistrict"

	mainCtx := context.Background()

	countryCtx := context.WithValue(mainCtx, countryKey, "Indonesia")
	provinceCtx := context.WithValue(countryCtx, provinceKey, "Jawa Barat")
	districtCtx := context.WithValue(provinceCtx, districtKey, "Cianjur")
	subdistrictCtx := context.WithValue(districtCtx, subdistrictKey, "Karangtengah")

	fmt.Println("Country", countryCtx)
	fmt.Println("Province", provinceCtx)
	fmt.Println("District", districtCtx)
	fmt.Println("Sub District", subdistrictCtx)

	fmt.Println("========================================")

	fmt.Println("Country", countryCtx.Value("country"))
	fmt.Println("Province", provinceCtx.Value("province"))
	fmt.Println("District", districtCtx.Value("district"))
	fmt.Println("Sub District", subdistrictCtx.Value("subdistrict"))
}

func CreateCounter(ctx context.Context, wg *sync.WaitGroup, withTimeout bool) <-chan int {
	countChan := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(countChan)

		counter := 0

		for {
			select {
			case <-ctx.Done():
				return
			case countChan <- counter:
				counter++

				if withTimeout {
					time.Sleep(2 * time.Second)
				}
			}
		}
	}()

	return countChan
}

func TestContextWithCancel(t *testing.T) {
	mainCtx := context.Background()
	counterCtx, cancel := context.WithCancel(mainCtx)
	group := sync.WaitGroup{}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	total := CreateCounter(counterCtx, &group, false)

	for n := range total {
		fmt.Println("Current:", n)

		if n == 10 {
			break
		}
	}

	cancel()

	group.Wait()

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	mainCtx := context.Background()
	counterCtx, cancel := context.WithTimeout(mainCtx, 10*time.Second)
	group := sync.WaitGroup{}

	defer cancel()

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	total := CreateCounter(counterCtx, &group, true)

	for n := range total {
		fmt.Println("Current:", n)
	}

	group.Wait()

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
