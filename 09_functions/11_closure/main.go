package main

import "fmt"

func main() {
	var x int
	fmt.Println("x=", x)
	x++
	fmt.Println("x=", x)

	{
		y := 42
		fmt.Println("y=", y)
	}

	// fmt.Println("y=", y)   error, y is not in this scope

	foo()
}

func foo() {
	fmt.Println("Foo here")
	// fmt.Println("x=", x) error, x is not in this scope
}
