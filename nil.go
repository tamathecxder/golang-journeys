package main

import "fmt"

func renderName(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	name1 := renderName("Agus")
	name2 := renderName("")

	fmt.Println(name1)
	fmt.Println(name2)

	if name1 == nil {
		fmt.Println("Empty Name 1")
	}

	if name2 == nil {
		fmt.Println("Empty Name 2")
	}
}
