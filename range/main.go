package main

import "fmt"

type user struct {
	name    string
	age     int
	address string
}

func main() {
	jekonoMap := map[string]string{
		"first":  "one",
		"second": "two",
	}
	fmt.Println(jekonoMap)

	for key, value := range jekonoMap {
		fmt.Printf("key: %s, value: %s \n", key, value)
	}

	arr := [3]string{"one", "two", "three"}
	for index, value := range arr {
		fmt.Printf("index: %d, value: %s \n", index, value)
	}

	slice := []string{"one", "two", "three"}
	for index, value := range slice {
		fmt.Printf("index: %d, value: %s \n", index, value)
	}

	name := "Jekono" // []bite --> biteslice

	for index, char := range name {
		fmt.Printf("index: %d, char: %c \n", index, string(char))
	}
}
