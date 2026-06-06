package main

import "fmt"

// callback function in go
func calculate(x, y int, operation func(x int, y int) int) int {
	return operation(x, y)
}

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(calculate(10, 20, add))

}
