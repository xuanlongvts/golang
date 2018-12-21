package main

import "fmt"

func main() {
	x := 10

	fmt.Println("main value x before = ", x)
	fmt.Println("main address x before = ", &x)

	foo(&x)

	fmt.Println("==============================")
	fmt.Println("main value x after = ", x)
	fmt.Println("main address x after = ", &x)
}

func foo(y *int) {
	fmt.Println("foo value y = ", *y)
	fmt.Println("foo address y = ", y)

	*y = 1000
	fmt.Println("foo value y after = ", *y)
	fmt.Println("foo address y after = ", y)
}
