package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("I Love Parami", i)
	}

	for j := 0; j < 10; j++ {
		if j == 5 {
			break // stop loop
		}
		fmt.Println("I Love Parami1", j)
	}
	for k := 0; k < 10; k++ {
		if k%2 != 0 {

			continue // skip the iteration, go to next iteration
		}
		fmt.Println("I Love Parami1", k)
	}

	limit := 1
	for limit <= 5 {
		fmt.Println("Jekono print:", limit)
		limit++
	}
}
