package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Main here")
}

func foo() {
	defer func() {
		fmt.Println("Defer in Foo")
	}()
	fmt.Println("Foo here")
}
