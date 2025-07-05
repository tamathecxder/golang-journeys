package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestGreetingAsep(t *testing.T) {
// 	result := Greeting("Asep")

// 	if result != "Hello, Asep" {
// 		t.Error("Result is not 'Hello, Asep'")
// 	}

// 	fmt.Println("TestGreetingAsep is done.")
// }

// func TestGreetingJohn(t *testing.T) {
// 	result := Greeting("John")

// 	if result != "Hello, John" {
// 		t.Fatal("Result is not 'Hello, John'")
// 	}

// 	fmt.Println("TestGreetingJohn is done.")
// }

func TestGreetingAssert(t *testing.T) {
	assert.Equal(t, "Hello, xAsep", Greeting("Asep"), "The result must be 'Hello, Asep'")
	fmt.Println("TestGreetingAsep is done.")
}

func TestGreetingRequire(t *testing.T) {
	require.Equal(t, "Hello, xJohn", Greeting("John"), "The result must be 'Hello, John'")
	fmt.Println("TestGreetingJohn is done.")
}
