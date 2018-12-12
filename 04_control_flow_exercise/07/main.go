package main

import (
	"fmt"
)

func main() {
	name := "Long Le"

	if name == "Nong No" {
		fmt.Println("Nong No")
	} else if name == "Long Le" {
		fmt.Println("name = ", name)
	} else {
		fmt.Println("either")
	}
}
