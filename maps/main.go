package main

import "fmt"

type user struct {
	name    string
	age     int
	address string
}

func main() {
	myMap := make(map[string]int)

	myMap["one"] = 1
	myMap["two"] = 2

	jekonoMap := map[string]string{
		"first":  "one",
		"second": "two",
	}

	fmt.Println(myMap)
	fmt.Println("one:", myMap["one"])
	fmt.Println(jekonoMap)
	fmt.Println("first:", jekonoMap["first"])

	delete(jekonoMap, "first")

	userMap := map[string]user{
		"data": {
			name:    "Jekono",
			age:     30,
			address: "Bandarban",
		},
	}
	fmt.Println(userMap)
}
