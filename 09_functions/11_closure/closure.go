package main

import "fmt"

func main() {
	plus1 := incrementor()
	plus2 := incrementor()

	fmt.Println("plus1=", plus1())
	fmt.Println("plus1=", plus1())
	fmt.Println("plus1=", plus1())

	fmt.Println("=====================")
	fmt.Println("plus2=", plus2())
	fmt.Println("plus2=", plus2())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
