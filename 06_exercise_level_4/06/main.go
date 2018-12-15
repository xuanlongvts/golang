package main

import "fmt"

func main() {
	x := make([]string, 3, 6)
	x = []string{"Mot", "Hai", "Ba", "Bon", "Nam", "Sau"}
	fmt.Println("x = ", x)
	fmt.Println("length x = ", len(x))
	fmt.Println("capacity x = ", cap(x))

	getLengX := len(x)
	for i := 0; i < getLengX; i++ {
		fmt.Println("index, value = ", i, x[i])
	}
	fmt.Println("====================")
	for i, v := range x {
		fmt.Println("index, value = ", i, v)
	}
}
