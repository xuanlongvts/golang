package main

import "fmt"

func main() {
	m := map[string]int{
		"james":      32,
		"Miss money": 25,
	}

	fmt.Println("map = ", m)
	fmt.Println("map(james) = ", m["james"])
	fmt.Println("map(banana) = ", m["banana"])

	v, ok := m["banana"]
	fmt.Println("v = ", v)
	fmt.Println("ok = ", ok)

	fmt.Println("=====================")
	if v, condition := m["Miss money"]; condition {
		fmt.Println("if print ", v)
	}
}
