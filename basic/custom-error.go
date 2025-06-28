package basic

import "fmt"

type notFoundError struct {
	Message string
}

type validationError struct {
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func (e *validationError) Error() string {
	return e.Message
}

type User struct {
	Id, Name string
}

func GetUser(id string, user *User) (*User, error) {
	if id == "" {
		return nil, &validationError{Message: "ID is required."}
	}

	if user.Id != id {
		return nil, &notFoundError{Message: "User is not found."}
	}

	return user, nil
}

func Basic_CustomError() {
	mockUser := User{"1", "Admin"}

	user, err := GetUser("2", &mockUser)

	if err != nil {
		// if validationErr, status := err.(*validationError); status {
		// 	fmt.Println(validationErr.Error(), "Status:", status)
		// } else if internalError, status := err.(*notFoundError); status {
		// 	fmt.Println(internalError.Error(), "Status:", status)
		// } else {
		// 	fmt.Println("Something went wrong, please try again later")
		// }

		switch err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", err.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", err.Error())
		default:
			fmt.Println("Something went wrong, please try again later")
		}
	} else {
		fmt.Println("User:", user)
	}

}
