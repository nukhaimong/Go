package main

import "fmt"

// variadic function = a function that can take a variable number of arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func greet(prefix string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s %s \n", prefix, name)
	}
}

func main() {
	sum := sum(10, 20, 30, 40, 50)
	fmt.Println(sum)
	people := []string{"Alice", "Bob", "Charlie"}
	greet("Hello,", "Alice", "Bob", "Charlie")
	greet("Hi,", people...) // variadic argument
}
