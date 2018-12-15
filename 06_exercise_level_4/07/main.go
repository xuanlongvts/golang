package main

import "fmt"

func main() {
	x := []string{"Mot", "Hai", "Ba"}
	y := []string{"Bon", "Nam", "Sau"}

	xy := [][]string{x, y}
	fmt.Println("xy = ", xy)

	for i, v := range xy {
		fmt.Println("record: ", i)
		for ii, vv := range v {
			fmt.Printf("\t index position: %v \t value: %v \n", ii, vv)
		}
	}
}
