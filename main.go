package main

import "fmt"

func main() {
	var txt string
	txt = "Hello, world"
	var firstName string = "Jekono Marma"
	lastName := "Nu Khai" // most used in real world application, and it cannot be used outside a function
	var age = 55

	// multiple variables declaration
	var (
		fatherName string = "Sr. Jekono"
		fatherAge  int    = 60
	)
	var x, y int
	x = 20
	y = 30

	var t, f int = 25, 26

	// constant variable declaration
	const pi = 3.14

	fmt.Println(txt, ",", firstName, lastName, "my age is", age, fatherName, fatherAge)
	fmt.Println(x, y, t, f)

	// use case of printf
	course := "NextLevel"
	batch := 6
	rating := 4.5

	fmt.Printf("This is %s course of web dev, batch %d and the course rating is %f", course, batch, rating)

	formattedString := fmt.Sprintf("This is %s course of web dev, batch %d and the course rating is %f", course, batch, rating)

	fmt.Println(formattedString)
}
