package main

import "fmt"

type Person struct {
	first string
}

type human interface {
	speak()
}

func (P *Person) speak() {
	fmt.Println("Hello:")
}

func saySomething(H human) {
	H.speak()
}

func main() {
	p1 := Person{
		first: "Long",
	}
	saySomething(&p1)

	fmt.Println("===============")

	p1.speak()
}
