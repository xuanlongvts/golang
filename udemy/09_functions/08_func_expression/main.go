package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Func expression no params")
	}
	f()

	g := func(x int) {
		fmt.Println("Func expression with param x=", x)
	}
	g(100)
}
