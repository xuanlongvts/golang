package main

import "fmt"

func main() {
	f := foo()

	fmt.Println("f=", f())
}

func foo() func() int {
	return func() int {
		return 40
	}
}
