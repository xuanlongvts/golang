package main

import "fmt"

func main() {
	xi := []int{2, 4, 6, 8}
	x := sum(xi...)
	fmt.Println("sum = ", x)

	y, z := foo("Long") // miss x ...int
	fmt.Println("y=", y)
	fmt.Println("z=", z)
	fmt.Println("========================")
	g, h := foo("Le", 2, 3, 4, 5)
	fmt.Println("g=", g)
	fmt.Println("h=", h)
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

func foo(s string, x ...int) (string, int) {
	sum := 0
	for _, v := range x {
		sum += v
	}

	str := "Hello, " + s

	return str, sum
}
