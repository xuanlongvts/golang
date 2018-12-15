package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 9, 11, 13, 15}

	fmt.Println("x = ", x)
	fmt.Println("(index) 0 < 3 = ", x[:3])
	fmt.Println("(index) >=3 = ", x[3:])
	fmt.Println("(index) from 2 to < 5 = ", x[2:5])
}
