package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Long",
		Last:  "Le",
		Age:   30,
	}
	p2 := Person{
		First: "Nong",
		Last:  "No",
		Age:   20,
	}
	person := []Person{p1, p2}
	fmt.Println("person =", person)
	fmt.Println("=======================")
	bs, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error message:", err)
	}
	// fmt.Println(bs)
	fmt.Println(string(bs))
}
