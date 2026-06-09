package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println(result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divided by zero")
	}
	return a / b, nil
}
