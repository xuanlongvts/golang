package main

import (
	"fmt"
)

func main() {
	x := []int{1, 3, 5, 7}

	fmt.Println("length len(x) = ", len(x))
	fmt.Println("length cap(x) = ", cap(x))
	fmt.Println("x = ", x)
	fmt.Println("x1 = ", x[1])

	for index, value := range x {
		fmt.Println("index, value = ", index, value)
	}
}
