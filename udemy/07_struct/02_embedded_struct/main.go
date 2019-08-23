package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Long",
			last:  "Le",
			age:   20,
		},
		first: "Something is out person",
		ltk:   true,
	}

	fmt.Println("sa1 = ", sa1)

	fmt.Println("==============================")
	fmt.Println("sa1: ", sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println("or: ", sa1.person.first, sa1.first)
}
