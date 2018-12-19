package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := sum(number...)

	fmt.Println("result=", result)

	sumEven := even(sum, number...)
	fmt.Println("sumEven=", sumEven)
}

func sum(xi ...int) int {
	var total = 0
	for _, v := range xi {
		total += v
	}

	return total
}

func even(f func(xi ...int) int, vi ...int) int { // tinh tong so chan
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
