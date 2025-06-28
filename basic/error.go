package basic

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide the number by 0")
	}

	return x / y, nil
}

func Basic_Error() {
	res, err := divide(20, 5)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
