package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 2, 5, 99, 22, 33, 34, 21, 98, 222, 96}
	xs := []string{"Chin", "Tam", "Sau muoi ba", "Muoi lam"}

	fmt.Println("xi1 =", xi)
	sort.Ints(xi)
	fmt.Println("xi2 =", xi)

	fmt.Println("==================================")
	fmt.Println("xs1 =", xs)
	sort.Strings(xs)
	fmt.Println("xs2 =", xs)
}
