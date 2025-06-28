package stdlib

import (
	"errors"
	"fmt"
	"strings"
)

var (
	validationError     error = errors.New("Validation Error Occcured")
	internalServerError error = errors.New("Internal Server Error Occcured")
)

const IS_MAINTENANCE bool = true

func GetAdminByEmail(email string) error {
	if email == "" {
		return validationError
	}

	if strings.Contains(email, "@gmail.com") {
		return validationError
	}

	if email != "admin@go.com" {
		return internalServerError
	}

	return nil
}

func Stdlib_Errors() {
	email := "admin@go.com"
	err := GetAdminByEmail(email)

	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("There are some required or invalid fields")
		} else if errors.Is(err, internalServerError) {
			fmt.Println("Oops! we ran out of trouble!")
		} else {
			fmt.Println("Unkown")
		}
	} else {
		fmt.Println("Success")
	}
}
