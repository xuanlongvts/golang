package main

import "fmt"

func main() {
	x := 42
	fmt.Println("x1 = ", x)
	x = 99
	fmt.Println("x2 = ", x)
	y := 100 + 70
	fmt.Println("y = ", y)
}

// statement is x := 42 or fmt.Println("x1 = ", x)
// expression is 100 + 70
// operand is 100 or 70
// operation is :=
