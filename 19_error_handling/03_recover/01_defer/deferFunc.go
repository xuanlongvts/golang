package main

import "fmt"

func main() {
	var x int
	x++
	fmt.Println("x=", x)

	z := c()
	fmt.Println("z=", z)
}

func c() (i int) {
	defer func() {
		i++
	}()

	return 5
}
