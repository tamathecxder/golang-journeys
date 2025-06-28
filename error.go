package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide the number by 0")
	}

	return x / y, nil
}

func main() {
	res, err := divide(20, 5)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
