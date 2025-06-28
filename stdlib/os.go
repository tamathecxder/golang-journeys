package stdlib

import (
	"fmt"
	"os"
)

func Stdlib_OS() {
	hn, err := os.Hostname()

	if err != nil {
		fmt.Println("Hostname", hn)
	}

	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}
}
