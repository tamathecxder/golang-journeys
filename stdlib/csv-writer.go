package stdlib

import (
	"encoding/csv"
	"os"
)

func Stdlib_CsvWriter() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}
