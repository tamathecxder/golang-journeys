package main

import "fmt"

func whatever() any { // or interface{}
	return 123
	// return false
	// return "whaterver"
}

func main() {
	fmt.Println(whatever())
}
