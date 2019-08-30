package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10) // limit core CPU

	numberProcessor := runtime.NumCPU()
	fmt.Println("Number CPU = ", numberProcessor)

	numberP1 := runtime.GOMAXPROCS(0)
	fmt.Println("Nubmer P1 = ", numberP1)
}