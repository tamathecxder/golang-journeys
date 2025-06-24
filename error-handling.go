package main

import "fmt"

func methodWithIssue(isError bool) {
	defer errorHandling()

	if isError {
		panic("Something went wrong!")
	}
}

func errorHandling() {
	msg := recover()
	fmt.Println("Error:", msg)
}

func logging() {
	fmt.Println("Call Logging...")
}

func initApp() {
	fmt.Println("Init")

	defer logging()
}

func main() {
	initApp()

	methodWithIssue(true)
}
