package basic

import (
	"fmt"
	"golang-journeys/basic/database"
	_ "golang-journeys/basic/orm"
)

func Basic_PackageInit() {
	driver := database.GetDBDriver()

	fmt.Println(driver)
}
