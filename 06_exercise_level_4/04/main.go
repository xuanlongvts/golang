package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 9}
	fmt.Println("x = ", x)
	x = append(x, 10)
	fmt.Println("x = ", x)
	var y = []int{15, 20, 25, 30}
	fmt.Println("y = ", y)
	xy := append(x, y...)
	fmt.Println("xy = ", xy)
}
