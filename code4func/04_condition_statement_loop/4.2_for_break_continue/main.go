package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if (i == 4) {
			break
		}
		fmt.Print("i = ", i, ", ")
	}

	fmt.Println("")
	fmt.Println("Out for")

	fmt.Println("==============")
	for i := 0; i < 10; i++ {
		if (i == 4) {
			continue
		}
		fmt.Print("i = ", i, ", ")
	}

	fmt.Println("")
	fmt.Println("==============")
	// while
	j := 0
	for j < 10 {
		fmt.Print("j = ", j, ", ")
		j++
	}
}