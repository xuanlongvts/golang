package main

import "fmt"

func main() {
	fmt.Println("Hellow world")

	foo()

	fmt.Println("Hellow Nong No")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("i: ", i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm food")
}

func bar() {
	fmt.Println("Bar is mine")
}
