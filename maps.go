package main

import "fmt"

func main() {
	biodata := map[string]string{
		"name": "Agus",
		"age":  "20",
	}

	biodata["name"] = "Gilang"
	biodata["gaming"] = "xxx"

	fmt.Println(biodata)
	fmt.Println(biodata["name"])
	fmt.Println(biodata["age"])
	fmt.Println(biodata["gaming"])

	delete(biodata, "gaming")

	fmt.Println(biodata)

	address := make(map[string]string)

	address["long"] = "6.2313123"
	address["lat"] = "12.21314"

	fmt.Println(address)
}
