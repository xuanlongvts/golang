package main

import "fmt"

var x string
var y int

func main() {
	// Declare a variable to be of a certain TYPE and then ASIGN a VALUE of that type to the variable
	fmt.Println("x = ", x)
	fmt.Printf("%T\n", x)

	fmt.Println("y = ", y)
	fmt.Printf("%T\n ", y)
}

/*
	understanding ZERO value
		false for booleans
		0 for integers
		0.0 for float
		"" for string

		nil for
			pointers
			functions
			interfaces
			slices
			channels
			maps
*/
