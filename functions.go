package main

import "fmt"

func saySalam(akhi string) (string, bool) {
	return "Assalamualaikum " + akhi, akhi != ""
}

func summa(a, b int) (total int) {
	total = a + b

	return total
}

func main() {
	akhi1, _ := saySalam("")

	fmt.Println(akhi1)

	result := summa(20, 40)

	fmt.Println(result)
}
