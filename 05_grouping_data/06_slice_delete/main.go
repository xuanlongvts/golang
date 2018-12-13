package main

import "fmt"

func main() {
	var x = []int{1, 3, 5, 7, 9, 11, 15, 17, 19}
	x = append(x[:2], x[4:]...)
	fmt.Println("after delete x = ", x)
}
