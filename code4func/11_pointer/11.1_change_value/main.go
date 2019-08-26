package main

import "fmt"

func main() {

	// pointer -> number
	var numberOne int = 100
	var pointerOne *int
	pointerOne = &numberOne

	fmt.Println("numberOne 1 = ", numberOne)
	*pointerOne = 999
	fmt.Println("numberOne 2 = ", numberOne)

	// pointer -> array
	var arrA = [3]int {1, 2, 3}
	var pointerTwo *[3]int
	pointerTwo = &arrA
	fmt.Println("pointer two = ", pointerTwo)
	fmt.Printf("type of pointer array: %T", pointerTwo)

	// call function pointer
	fmt.Println("")
	fmt.Println("========================")
	var numberTwo = 555
	var pointerThree *int
	pointerThree = &numberTwo
	applyPointer(pointerThree)
	fmt.Println("value numberTwo = ", numberTwo)
}

func applyPointer(pointer *int) {
	*pointer = 777
}