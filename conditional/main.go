package main

import "fmt"

// if else
func grade() {

	//score := 20
	//if scope

	if score := 20; score >= 80 {
		fmt.Println("You got A+, your score is:", score)
	} else if score < 80 && score >= 40 {
		fmt.Println("you're pass, your score is:", score)
	} else {
		fmt.Println("You're fail, your score is:", score)
	}

}

// switch
func calDay(day string) {
	switch day { // tag switch
	case "Saturday":
		fmt.Println("Off day")
	case "Friday":
		fmt.Println("weekend")
	default:
		fmt.Println("wroking day")
	}
	switch { // normal switch
	case day == "Saturday":
		fmt.Println("Off day")
	case day == "Friday":
		fmt.Println("weekend")
	default:
		fmt.Println("wroking day")
	}
}

func main() {
	grade()
	calDay("Teusday")
}
