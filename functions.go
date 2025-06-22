package main

import "fmt"

func saySalam(akhi string) (string, bool) {
	return "Assalamualaikum " + akhi, akhi != ""
}

func summa(a, b int) (total int) {
	total = a + b

	return total
}

func decreaseAll(nums ...int) int {
	total := 0
	for _, num := range nums {
		total -= num
	}

	return total
}

type Person map[string]string

func injectPerson(name string, age string) Person {
	person := make(Person)
	person["name"] = name
	person["age"] = age

	return person
}

func main() {
	akhi1, _ := saySalam("")

	fmt.Println(akhi1)

	result := summa(20, 40)

	fmt.Println(result)

	nums := []int{20, 10, 30, 40}

	fmt.Println(decreaseAll(20, 20, 50, 50))
	fmt.Println(decreaseAll(nums...))

	person := injectPerson

	fmt.Println(person("Ujang", "22"))
}
