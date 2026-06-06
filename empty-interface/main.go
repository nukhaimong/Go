package main

import "fmt"

// any == empty interface == interface{}
// type assertion
// ok idiom

func print(data any) {
	strData, ok := data.(string)
	if !ok {
		fmt.Println("Data is not a string")
		return
	}
	fmt.Println(strData)
}

func main() {
	print(10)
}
