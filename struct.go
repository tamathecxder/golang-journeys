package main

import (
	"fmt"
	"math/rand"
)

type User interface {
	GetActivityStatus() string
}

type Employee struct {
	Id, Name    string
	BasicSalary int
}

func (e Employee) GetActivityStatus() string {
	return "Online"
}

func (e Employee) GetSalary() string {
	return fmt.Sprintf("Rp. %d", e.BasicSalary)
}

func generateID() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const maxLength = 10

	id := make([]byte, maxLength)

	for i := 0; i < maxLength; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}

	return string(id)
}

func main() {
	var e1 Employee

	e1.Id = generateID()
	e1.Name = "Agus"
	e1.BasicSalary = 2500000

	fmt.Println(e1)

	e2 := Employee{
		Id:          generateID(),
		Name:        "Asep",
		BasicSalary: 20000,
	}

	fmt.Println(e2.Id, e2.Name, e2.BasicSalary)

	e3 := Employee{generateID(), "Ahmad", 290000}

	fmt.Println(e3.GetActivityStatus())
	fmt.Println(e3.GetSalary())
}
