package main

import "fmt"

// data struct can composing different type data together
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Long",
		last:  "Le",
		age:   20,
	}

	p2 := person{
		first: "Nong",
		last:  "No",
		age:   30,
	}

	fmt.Println("person1 = ", p1)
	fmt.Println("person2 = ", p2)

	fmt.Println("===============")

	fmt.Println("1: ", p1.first, p1.last, p1.age)
	fmt.Println("2: ", p2.first, p2.last, p2.age)
}
