package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous function")
	}()

	func(x int) {
		fmt.Println("Anonymous function x=", x)
	}(42)
}

func foo() {
	fmt.Println("Foo here")
}
