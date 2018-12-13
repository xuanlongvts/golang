package main

import "fmt"

func main() {
	m := map[string]int{
		"mot": 1,
		"hai": 2,
		"ba":  3,
	}
	fmt.Println("m = ", m)

	delete(m, "hai")
	fmt.Println("m = ", m)

	if v, ok := m["ba"]; ok {
		fmt.Println("map key ba = ", v)
		delete(m, "ba")
	}

	fmt.Println("============")
	fmt.Println("m = ", m)
}
