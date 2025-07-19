package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	totalGmc := runtime.GOMAXPROCS(-1)
	fmt.Println("Total GMC:", totalGmc)

	totalGo := runtime.NumGoroutine()
	fmt.Println("Total GO:", totalGo)

}
