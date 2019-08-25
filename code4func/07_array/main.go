package main

import "fmt"

func main(){

	// declation
	var arrA [4]int

	fmt.Println("arrA 1 = ", arrA)

	arrA[0] = 123
	fmt.Println("arrA 2 = ", arrA)

	arrA[3] = 456
	fmt.Println("arrA 2 = ", arrA)

	// declation and init value
	fmt.Println("=================")
	arrB := [3]int {1, 2, 3}
	fmt.Println("arrB = ", arrB)

	// declation and init not full value
	fmt.Println("=================")
	arrC := [3]int {100}
	fmt.Println("arrC = ", arrC, ", length = ", len(arrC))

	// declation but not set size, automatic count size array
	fmt.Println("=================")
	arrD := [...]string {"mot", "hai", "ba", "bon", "nam"}
	fmt.Println("arrD = ", arrD, ", length = ", len(arrD), ", value at index 2 = ", arrD[2])

	// array is value type not reference type
	fmt.Println("=================")
	arrCountryOne := [...]string {"VietNam", "US", "France"}

	arrCountryTwo := arrCountryOne
	arrCountryTwo[0] = "Canada"
	fmt.Println("arrCountryOne = ", arrCountryOne, ", arrCountryTwo = ", arrCountryTwo)
}