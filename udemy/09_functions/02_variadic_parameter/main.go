package main

import "fmt"

func main() {
	x := sum(1, 3, 5, 7) // can pass none parameter x := sum(), variadic from zero to more
	fmt.Println("sum = ", x)
}

func sum(x ...int) int {
	fmt.Println("x=", x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
