package main

import "fmt"

func main() {
	m := map[string]int{
		"mot": 1,
		"hai": 2,
	}
	fmt.Println("m = ", m)

	m["ba"] = 3
	fmt.Println("m = ", m)

	for k, v := range m {
		fmt.Println("k=", k, ", v = ", v)
	}

	fmt.Println("===================")
	x := []int{1, 3, 5, 7}
	for k, v := range x {
		fmt.Println("k=", k, ", v = ", v)
	}
}
