package main

import (
	"fmt"
)

func main() {
	year := 2000
	for year < 2005 {
		fmt.Println("year = ", year)

		year++
	}
}
