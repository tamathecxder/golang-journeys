package stdlib

import (
	"bufio"
	"os"
)

func Stdlib_BufioWriter() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Assalamualaikum\n")
	_, _ = writer.WriteString("Sohabat\n")

	writer.Flush()
}
