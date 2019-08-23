package main

import "fmt"

func main() {
	arg := []int{1, 2, 3, 4, 5}
	result := sum(arg...)
	fmt.Println("result=", result)

	arg2 := []int{2, 4, 6}
	result2 := bar(arg2)
	fmt.Println("result2=", result2)
}

func sum(xi ...int) int {
	var total = 0
	for _, v := range xi {
		total += v
	}

	return total
}

func bar(xi []int) int {
	var total = 0
	for _, v := range xi {
		total += v
	}

	return total
}
