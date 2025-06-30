package stdlib

import (
	"fmt"
	"path"
)

func Stdlib_Path() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}
