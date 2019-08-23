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

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "-- secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "-- person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
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
	sa1.speak()
	sa2.speak()

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println("p1=", p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
