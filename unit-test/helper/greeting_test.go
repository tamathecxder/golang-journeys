package helper

import (
	"fmt"
	"testing"
)

func TestGreetingAsep(t *testing.T) {
	result := Greeting("Asep")

	if result != "Hello, Asep" {
		t.Error("Result is not 'Hello, Asep'")
	}

	fmt.Println("TestGreetingAsep is done.")
}

func TestGreetingJohn(t *testing.T) {
	result := Greeting("John")

	if result != "Hello, John" {
		t.Fatal("Result is not 'Hello, John'")
	}

	fmt.Println("TestGreetingJohn is done.")
}
