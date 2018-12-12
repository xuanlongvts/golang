package main

import (
	"fmt"
)

func main() {
	year := 2000
	for year < 2008 {
		if year > 2005 {
			break
		}
		fmt.Println("year = ", year)
		year++
	}
}
