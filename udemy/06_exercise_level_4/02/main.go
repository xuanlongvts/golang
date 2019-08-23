package main

import "fmt"

func main() {
	// Slice type
	var x = []int{1, 3, 5, 7, 9, 11}
	for i, v := range x {
		fmt.Println("index=", i, ", value=", v)
	}

	fmt.Printf("%T\n", x)
}
