package main

import "fmt"

type Ghost struct {
	Name, Origin string
}

func main() {
	var gh1 *Ghost = new(Ghost) // new method can be using to create a pointer, but the declarations can be empty or not initialize at first
	var gh2 *Ghost = gh1

	gh2.Name = "Pocong"

	fmt.Println(gh1)
	fmt.Println(gh2)
}
