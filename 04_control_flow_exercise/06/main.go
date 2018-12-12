package main

import (
	"fmt"
)

func main() {
	for i := 10; i < 100; i++ {
		fmt.Printf("When %v divided 4, remainer = %v \n", i, i%4)
	}
}
