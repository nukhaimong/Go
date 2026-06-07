package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"-"` // this field will be ignored during JSON encoding and decoding
	City string `json:"city"`
}

func main() {
	p := person{
		Name: "Jekono",
		Age:  30,
		City: "Lagos",
	}
	rawJson, err := json.Marshal(p)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(string(rawJson))

	var p2 person
	err = json.Unmarshal(rawJson, &p2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p2)
}
