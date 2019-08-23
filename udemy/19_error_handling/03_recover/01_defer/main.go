package main

import "fmt"

func main() {
	i := 0

	defer fmt.Println("i=", i) // i = 0 , if we put after i++, i will equal = 1

	i++
	return
}
