package  main

import "fmt"

func main(){
	// declation slice
	var slice []int
	fmt.Println("slice = ", slice)

	// declation slice and init
	var sliceOne = []int {1, 2, 3, 4}
	fmt.Println("slice one = ", sliceOne)

	// create a slice from an array
	var arrA = [8]int {1, 2, 3, 4, 5, 6, 7, 8}
	var sliceTwo = arrA[1:6] // arrA[1] -> arrA[6-1]
	fmt.Println("slice two = ", sliceTwo)

	var sliceThree = arrA[:]
	fmt.Println("slice three = ", sliceThree)

	var sliceFour = arrA[2:] // arrA[2] -> end
	fmt.Println("slice four = ", sliceFour)

	var sliceFive = arrA[:3] // arrA[0] -> arrA[3 - 1]
	fmt.Println("slice five = ", sliceFive)

	// slice with slice
	// create a slice from other slice
	var sliceSix = []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceSeven := sliceSix
	fmt.Println("Slice seven = ", sliceSeven)

	sliceEight := sliceSix[2:]
	fmt.Println("Slice eight = ", sliceEight)
}