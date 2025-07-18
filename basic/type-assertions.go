package basic

import "fmt"

func anything() any {
	return "OK"
}

func Basic_TypeAssertions() {
	data := anything()
	fmt.Println(data.(string)) // manual

	// fmt.Println(data.(int)) // panic

	switch value := data.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("default", value)
	}
}
