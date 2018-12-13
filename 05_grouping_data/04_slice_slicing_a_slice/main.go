package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 10}
	fmt.Println(x[1])
	fmt.Println("x = ", x)
	fmt.Println("x from index1 to end = ", x[1:])
	fmt.Println("x from index1 to index3 = ", x[1:3])

	for i, v := range x {
		fmt.Println("i, v ", i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println("value = ", x[i])
	}
}
