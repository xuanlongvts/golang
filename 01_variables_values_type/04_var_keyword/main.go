package main

import (
	"fmt"
)

// Declare z
var z = 21

// Declare VARIABLE y with identifier "y" of TYPE int and assign the ZERO value of TYPE int to "y"
var y int

func main() {
	// short declaration operator, declare and assign value(7) to x
	x := 7
	fmt.Println(x)
	foo()
	fmt.Println(y)
}

func foo() {
	fmt.Println(z)
}
