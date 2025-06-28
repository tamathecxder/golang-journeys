package basic

import "fmt"

func Basic_Operators() {
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
