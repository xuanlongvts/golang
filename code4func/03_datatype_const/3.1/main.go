package main

import "fmt"

func main(){
	// float
	var floatNumber float32 = 1.1
	fmt.Println("Number float = ", floatNumber)

	// complex
	var complexNumberOne complex64 = 10 + 1i
	fmt.Println("Number complex 1 = ", complexNumberOne)
	var complexNumberTwo complex64 = 10 + 2i
	fmt.Println("Number complex 2 = ", complexNumberTwo)
	var totalComplex = complexNumberOne + complexNumberTwo
	fmt.Println("Total complex = ", totalComplex)
}