package main

import "fmt"

func main() {
	s := foo()
	fmt.Println("string=", s)

	barVar := bar()
	fmt.Printf("%T\n", barVar)
	fmt.Println("value func= ", barVar())

	fmt.Println("call func direct = ", bar()())
}

func foo() string {
	s := "Hello world"

	return s
}

func bar() func() int {
	return func() int {
		return 456
	}
}
