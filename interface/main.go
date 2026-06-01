package main

import "fmt"

// interface = a type that specifies a set of method signatures (behavior)

type Dog struct{}
type Cat struct{}

type Animal interface {
	speak()
}

func (d Dog) speak() {
	fmt.Println("Woof!")
}

func (c Cat) speak() {
	fmt.Println("Meow!")
}

func makeSound(a Animal) {
	a.speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	makeSound(dog)
	makeSound(cat)
}
