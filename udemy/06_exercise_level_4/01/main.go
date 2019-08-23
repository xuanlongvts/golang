package main

import "fmt"

func main() {
	// Size array is 5, if we declare miss element then default is 0

	x := [5]int{1, 3, 5, 7}
	for i, v := range x {
		fmt.Println("i = ", i, ", value=", v)
	}
	fmt.Printf("%T\n", x)
}
