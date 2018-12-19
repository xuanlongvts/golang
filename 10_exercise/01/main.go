package main

import "fmt"

func main() {
	n := foo()
	x, s := bar()

	fmt.Println("n=", n)
	fmt.Println("x=", x)
	fmt.Println("s=", s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 100, "Hello world"
}
