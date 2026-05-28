package main

import "fmt"

func main() {

	var orders = [6]int{1, 3, 4, 5, 6, 6}
	slice := orders[1:4] // [start: end(excluded)]

	slice = append(slice, 200)
	fmt.Println(slice)

	fmt.Println("The length of the slice:", len(slice))
	fmt.Println("The capacity of the slice:", cap(slice))

	// direct declaration of slice
	var str = []string{"jekono", "nukhai"}
	fmt.Println(str)
	str = append(str, "Go")
	fmt.Println(cap(str))

}
