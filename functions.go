package main

import "fmt"

func saySalam(akhi string) (string, bool) {
	return "Assalamualaikum " + akhi, akhi != ""
}

func main() {
	akhi1, _ := saySalam("")

	fmt.Println(akhi1)
}
