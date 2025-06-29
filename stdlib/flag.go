package stdlib

import (
	"flag"
	"fmt"
)

func Stdlib_Flag() {
	var username *string = flag.String("username", "admin", "Username for login credentials")
	var password *string = flag.String("password", "passwor123", "Username for login credentials")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
}
