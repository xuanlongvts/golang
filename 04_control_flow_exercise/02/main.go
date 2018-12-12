package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 92; i++ {
		fmt.Println("i = ", i)
		for ii := 0; ii < 3; ii++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
