package stdlib

import (
	"fmt"
	"strings"
)

func Stdlib_Strings() {
	var names = strings.Split("Mamat_Gunshop", "_")
	for _, n := range names {
		fmt.Println(n)
	}

	fmt.Println(strings.Contains("Kantor", "or"))
	fmt.Println(strings.ToUpper("bunted"))
	fmt.Println(strings.ToLower("Hidup Jokowi!"))
	fmt.Println(strings.Trim(" JAYAJAYA JAYA!   ", " "))
	fmt.Println(strings.ReplaceAll("Jawa Timur", "Timur", "Barat"))
}
