package main

import "fmt"

// No pointer
func main() {
	x := 10
	foo(x)
	fmt.Println("main x = ", x)
}

func foo(y int) {
	fmt.Println("foo y1 = ", y)
	y = 40
	fmt.Println("foo y2 = ", y)
}
