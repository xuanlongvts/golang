package main

import "fmt"

func main() {
	fmt.Println("Hello golang")

	fmt.Println("Total = ", SumNumber(5))
}

func SumNumber(number int) int {
	sum := 0

	for i := 0; i < number; i++ {
		sum += i
	}

	return sum
}
