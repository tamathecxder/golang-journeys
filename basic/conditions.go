package basic

import "fmt"

func Basic_Conditions() {
	if true {
		fmt.Println("fr")
	} else {
		fmt.Println("ts pmo")
	}

	num := 0
	if isMore := num > 0; isMore {
		fmt.Println("ez")
	} else {
		fmt.Println("sybau")
	}

	name := "Agus"
	switch name {
	case "Agus":
		fmt.Println("ahay")
	default:
		fmt.Println("kys")
	}

	switch {
	case name == "":
		fmt.Println("xxx")
	default:
		fmt.Println("sds")
	}
}
