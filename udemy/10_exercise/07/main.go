package main

import "fmt"

var number func() = func() {
	fmt.Println("Function number outside")
}

func main() {

	f := func() {
		fmt.Println("Content in function")
	}
	f()
	fmt.Printf("%T\n", f)

	number()
	number = f
	fmt.Printf("%T\n", number)
	number()
}
