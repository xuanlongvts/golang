package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Long",
			last:  "Le",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Nong",
			last:  "No",
		},
		ltk: false,
	}
	fmt.Println("sa1=", sa1)
	sa1.speak()
	sa2.speak()
}
