package main

import "fmt"

var x int

type hotdog int

var y hotdog

func main() {
	x = 3
	y = 5
	fmt.Println("x = ", x)
	fmt.Printf("%T\n\n", x)

	fmt.Println("x = ", y)
	fmt.Printf("%T\n", y)

	// x = y cannot use y (type hotdog) as type int in assignment
}
