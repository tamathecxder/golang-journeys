package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("<BEFORE THE TEST>")

	m.Run()

	fmt.Println("<AFTER THE TEST>")
}

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

func TestGreetingAssert(t *testing.T) {
	assert.Equal(t, "Hello, xAsep", Greeting("Asep"), "The result must be 'Hello, Asep'")
	fmt.Println("TestGreetingAsep is done.")
}

func TestGreetingRequire(t *testing.T) {
	require.Equal(t, "Hello, xJohn", Greeting("John"), "The result must be 'Hello, John'")
	fmt.Println("TestGreetingJohn is done.")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run the test in MacOS")
	}

	assert.Equal(t, "Hello, xAsep", Greeting("Asep"), "The result must be 'Hello, Asep'")
	fmt.Println("TestGreetingAsep is done.")
}

func BenchmarkGreetings(b *testing.B) {
	b.Run("Agus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Greeting("Agus")
		}
	})

	b.Run("Suroso", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Greeting("Suroso")
		}
	})
}

func BenchmarkGreetingAgus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting("Agus")
	}
}

func BenchmarkGreetingAhmad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greeting("Ahmad")
	}
}
