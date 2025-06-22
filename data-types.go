package main

import "fmt"

func main() {
	// datatypes
	var isSundanese = true
	fmt.Println(isSundanese)

	var country string = "China"
	fmt.Println("Country = ", country)

	totalUsers := 0
	fmt.Println(totalUsers)

	trxPrice := 120.90
	fmt.Println(trxPrice)

	// datatype conversion
	var value32 int32 = 2221
	fmt.Println(value32)
	var value64 int64 = int64(value32)
	fmt.Println(value64)

	type postcode string

	var myPostCode postcode = "24364"
	fmt.Println(myPostCode)

	var firstNum = 20
	var secNum = 10
	var thirdNum = 0
	thirdNum += firstNum
	secNum++

	fmt.Println(firstNum * secNum)
	fmt.Println(thirdNum)
	fmt.Println(secNum)

	var isEqual = firstNum == secNum
	fmt.Println("is equal", isEqual)
}
