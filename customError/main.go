package main

import "fmt"

type CustomError struct {
	message string
	code    int
}

func (ce *CustomError) Error() string {
	return ce.message
}

func login(password string) error {
	if password != "secret" {
		return &CustomError{
			message: "Invalid password",
			code:    401,
		}
	}
	return nil
}

func main() {
	err := login("jekono")
	if err != nil {
		fmt.Println("Error", err, "Code:", err.(*CustomError).code)
	}
}
