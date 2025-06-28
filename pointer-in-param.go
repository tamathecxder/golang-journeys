package main

import "fmt"

type Phone struct {
	Name, Brand string
}

func changeBrandToSamsung(phone *Phone) {
	phone.Brand = "Samsung"
	fmt.Println(phone)
}

func main() {
	ph1 := Phone{}
	changeBrandToSamsung(&ph1) // it automatically turn this ph1 arg into pointer using & operator

	// var ph1 *Phone = new(Phone)
	// changeBrandToSamsung(ph1)

	fmt.Println(ph1)
}
