package main

import (
	"fmt"
	"golang-journeys/database"
	_ "golang-journeys/orm"
)

func main() {
	driver := database.GetDBDriver()

	fmt.Println(driver)
}
