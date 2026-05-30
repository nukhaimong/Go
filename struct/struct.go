package main

import "fmt"

type user struct {
	name     string
	age      int
	metaInfo additionalInfo
}

type additionalInfo struct {
	address    string
	profession string
}

func main() {
	user1 := user{
		name: "Jekono",
		age:  25,
		metaInfo: additionalInfo{
			address:    "123 Main St",
			profession: "Software Engineer",
		},
	}
	// fmt.Println(user1.name)
	// fmt.Printf("%+v", user1)

	user1.displayInfo()
	//pointerUser1 := &user1
	user1.updateAge()
	user1.displayInfo()
}

// reciever function
func (u user) displayInfo() {
	fmt.Printf("%+v", u)
}

// reciever function with pointer
func (u *user) updateAge() {
	//(*u).age = 30
	// alternative way to update age using pointer receiver
	u.age = 30
}
