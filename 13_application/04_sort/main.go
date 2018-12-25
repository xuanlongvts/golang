package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 1, 4, 0, 8, 9, 2, 7}
	xs := []string{"Mot", "Chin", "Q", "R", "K", "I", "New York", "Tam Quan Nam"}

	fmt.Println("before xi", xi)
	sort.Ints(xi)
	fmt.Println("sort xi", xi)
	fmt.Println("============================")
	fmt.Println("before xs =", xs)
	sort.Strings(xs)
	fmt.Println("sort xs=", xs)
}
