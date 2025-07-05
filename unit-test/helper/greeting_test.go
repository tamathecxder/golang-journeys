package helper

import "testing"

func TestGreeting(t *testing.T) {
	result := Greeting("Asep")

	if result != "Hello, Asep" {
		panic("Result is not 'Hello, Asep'")
	}
}
