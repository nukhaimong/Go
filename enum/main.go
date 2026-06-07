package main

import "fmt"

// enums
type weekDay int

const (
	sunday weekDay = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func getDays(day weekDay) {
	switch day {
	case monday, tuesday, wednesday, thursday, friday:
		fmt.Println("working day")
	case saturday:
		fmt.Println("half day")
	case sunday:
		fmt.Println("weekend")
	}
}

// enums in string
type weekDayString string

const (
	sundayString    weekDayString = "sunday"
	mondayString    weekDayString = "monday"
	tuesdayString   weekDayString = "tuesday"
	wednesdayString weekDayString = "wednesday"
	thursdayString  weekDayString = "thursday"
	fridayString    weekDayString = "friday"
	saturdayString  weekDayString = "saturday"
)

func main() {
	getDays(monday)
	getDays(saturday)
	getDays(sunday)
}
