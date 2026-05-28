package main

import "fmt"

// pointer in functions
func change(x *int) {
	*x = 2050
	fmt.Println("Inside change function, x:", *x) // prints 2050
	// note: pointers is memory efficient as it allows us to modify the original variable without creating a copy of it
}

func modifyArrWithPointer(arr *[5]int) {
	arr[0] = 100
}

func main() {
	// pointers
	a := 50
	p := &a // p is a pointer to a, storing the memory address of a

	a = 100
	*p = 200 // dereferencing p allows us to change the value of a through the pointer

	fmt.Println("a:", a)
	fmt.Println("p:", p)   // prints the memory address of a
	fmt.Println("*p:", *p) // dereferencing p gives the value of a, which is 100

	y := 10
	fmt.Println("Before change, y:", y) // prints 10

	change(&y) // passing the address of y to the change function

	fmt.Println("After change, y:", y) // prints 2050

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Before modifyArrWithPointer, arr:", arr) // prints [1 2 3 4 5]

	modifyArrWithPointer(&arr) // passing the address of the array to the function

	fmt.Println("After modifyArrWithPointer, arr:", arr) // prints [100 2 3 4 5]
}
