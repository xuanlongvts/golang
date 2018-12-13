package main

import "fmt"

func main() {
	x := []int{1, 3, 5}
	fmt.Println("x = ", x)
	x = append(x, 7, 9, 15)
	fmt.Println("x = ", x)

	y := []int{18, 20, 25}
	x = append(x, y...)
	fmt.Println("x = ", x)
}
