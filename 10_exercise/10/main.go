package main

import "fmt"

func main() {
	f := foo()
	fmt.Println("f=", f())
	fmt.Println("f=", f())
	fmt.Println("f=", f())
	fmt.Println("f=", f())

	fmt.Println("===============")
	g := foo()
	fmt.Println("g=", g())
	fmt.Println("g=", g())
	fmt.Println("g=", g())

}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
