package main

import "fmt"

// callback function in go
func calculate(x, y int, operation func(x int, y int) int) int {
	return operation(x, y)
}

// higher order function
func multiplyBy(factor int) func(int) int {
	return func(y int) int {
		return y * factor
	}
}

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(calculate(10, 20, add))

	// anonymouse callback function
	result := calculate(5, 10, func(x, y int) int {
		return x + y
	})

	fmt.Println(result)

	double := multiplyBy(2)
	fmt.Println(double(5))
}
