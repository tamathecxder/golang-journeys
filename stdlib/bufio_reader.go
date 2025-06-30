package stdlib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Stdlib_BufioReader() {
	input := strings.NewReader("Hello, this is your dad\nwelcome home son!\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
