package stdlib

import (
	"fmt"
	"math"
)

func Stdlib_Math() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Max(10, 11))
	fmt.Println(math.Min(10, 11))
}
