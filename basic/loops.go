package basic

import "fmt"

func Basic_Loops() {
	i := 1

	for i <= 10 {
		fmt.Println("I:", i)
		i++
	}

	status := []string{"Approved", "Rejected", "Pending"}
	for i, v := range status {
		fmt.Println("i", i, "v", v)
	}

	// golang support break and continue
}
