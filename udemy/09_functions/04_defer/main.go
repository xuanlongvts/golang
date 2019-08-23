package main

import "fmt"

func main() {
	defer foo()
	bar()
	drink()
}

func foo() {
	fmt.Println("Here Foo")
}

func bar() {
	fmt.Println("Bar here")
}

func drink() {
	fmt.Println("What do you drink?")
}
