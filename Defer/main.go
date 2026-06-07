package main

import "fmt"

func deffred(result int) {
	fmt.Println("Defer called with result:", result)
}

func deferExample() {
	result := 10

	defer deffred(result)

	// defer eith clousure
	defer func() {
		fmt.Println("Defer with closure called with result:", result)
	}()

	result += 5

	fmt.Println("Function called with result:", result)

}

func main() {
	deferExample()
}
