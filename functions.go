package main

import (
	"fmt"
	"strings"
)

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

func formatIntoAsterisk(text string) string {
	length := len(text)

	var result string

	for i := 1; i <= length; i++ {
		result += "*"
	}

	firstLetter := string(text[0])
	decreasedResult := result[1:]
	result = fmt.Sprintf("%s%s", firstLetter, decreasedResult)

	return result
}

type TextFormat func(string) string

func processText(text string, format TextFormat) string {
	badWords := []string{
		"Fuck",
		"Bitch",
		"Asshole",
	}

	for _, word := range badWords {
		if text == word {
			return format(text)
		}
	}

	return text
}

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
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

	fmt.Println("=======")

	text1 := "Fuck"
	text2 := "Genius"

	formatResult1 := processText(text1, formatIntoAsterisk)
	formatResult2 := processText(text2, formatIntoAsterisk)

	fmt.Println(formatResult1)
	fmt.Println(formatResult2)

	anonFuncResult := processText("Bitch", func(s string) string {
		if s == "" {
			return ""
		}

		return string(s[0]) + strings.Repeat("*", len(s)-1)
	})

	fmt.Println(anonFuncResult)

	fmt.Println("=======")

	fmt.Println(factorial(10))
}
