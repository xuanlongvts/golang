package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println("x = ", x)
	x[3] = 10
	fmt.Println("x = ", x)
	fmt.Println("length of x = ", len(x))
}
