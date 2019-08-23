package main

import "fmt"

type Person struct {
	name string
}

func changeMe(P *Person) {
	P.name = "Nong No"

	// or (*P).name = "Still change"
}

func main() {
	per := Person{
		name: "ABCD ABCD",
	}
	fmt.Println("person =", per)

	changeMe(&per)

	fmt.Println("after person=", per)
}
