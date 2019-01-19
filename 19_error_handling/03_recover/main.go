package main

import "fmt"

func main() {
	f()
	fmt.Println("Main value")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in f() func.", r)
			fmt.Println("===============")
		}
	}()

	fmt.Println("Call g")
	g(0)
	fmt.Println("Return from g() func") //  not call
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("Err panic: %v", i))
	}
	defer fmt.Println("Defer in g() func", i)
	fmt.Println("Print in g() func", i)
	g(i + 1)
}
