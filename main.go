package main

import "fmt"

func main() {
	var txt string
	txt = "Hello, world"
	var firstName string = "Jekono Marma"
	lastName := "Nu Khai"
	var age = 55

	var (
		fatherName string = "Sr. Jekono"
		fatherAge  int    = 60
	)

	fmt.Println(txt, ",", firstName, lastName, "my age is", age, fatherName, fatherAge)
}
