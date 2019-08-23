package main

import "fmt"

func main() {
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "Long",
		last:  "Le",
		age:   20,
	}

	fmt.Println("person = ", p)
	fmt.Println("first: ", p.first)
	fmt.Println("last: ", p.last)
	fmt.Println("age: ", p.age)
}
