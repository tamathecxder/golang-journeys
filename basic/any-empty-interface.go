package basic

import "fmt"

func whatever() any { // or interface{}
	return 123
	// return false
	// return "whaterver"
}

func Basic_AnyEmptyInterface() {
	fmt.Println(whatever())
}
