package main

import "fmt"

func main() {
	var names = [2]string{
		"Asep",
		"Ahmad",
	}

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(len(names))

	var unpredictedNumbers = [...]int32{
		20,
		124,
		12412,
		3453531,
		23434,
	}

	fmt.Println(unpredictedNumbers)
	fmt.Println(unpredictedNumbers[0])
	unpredictedNumbers[0] = 10
	fmt.Println(unpredictedNumbers[0])
}
