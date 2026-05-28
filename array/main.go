package main

import "fmt"

func main() {
	// array
	var numbers [6]int
	numbers[0] = 42
	numbers[1] = 24
	numbers[5] = 9

	// another way to define array
	arr := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)

	str := []string{"jekono", "nukhai"} // this is a slice, not an array. It is a reference type, and it can grow and shrink in size.
	fmt.Println(str)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("number:", numbers[i])
	// }

}
