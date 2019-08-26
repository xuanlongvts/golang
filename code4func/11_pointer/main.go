package main

import "fmt"

func main() {
	var number = 100
	var pointerOne *int
	pointerOne = &number

	fmt.Println("value of pointer is a Hexa = ", pointerOne)
	fmt.Printf("Type of pointer = %T", pointerOne)

	fmt.Println("")
	fmt.Println("===========================")
	var numberTwo = 1.1
	var pointerTwo = new(float64) // <=> var pointerTwo *float64
	pointerTwo = &numberTwo
	fmt.Println("value of pointer is a Hexa = ", pointerTwo)
	fmt.Printf("Type of pointer = %T", pointerTwo)

	// dung * thi Zero value se la nil, dung new thi Golang se tao ra ma Hexa
	fmt.Println("")
	fmt.Println("===========================")
	var pointerThree *int
	var pointerFour = new(int)
	fmt.Println("pointerThree = ", pointerThree, ", pointerFour = ", pointerFour)
}