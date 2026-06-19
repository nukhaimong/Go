package main

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{
		Id:    1,
		Name:  "Nu Khai",
		Age:   25,
		Email: "nukhai@gmail.com",
	},
	{
		Id:    2,
		Name:  "Jekono",
		Age:   25,
		Email: "jekono@gmail.com",
	},
}
