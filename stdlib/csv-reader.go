package stdlib

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func Stdlib_CsvReader() {
	csvString := "eko,kurniawan,khannedy\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
