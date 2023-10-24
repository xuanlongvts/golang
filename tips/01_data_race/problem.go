package main

import "fmt"

/*
A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write
*/

func main() {
	var getNum = getNumber()
	fmt.Println(getNum)
}

func getNumber() int {
	var i int
	go func() {
		i++
	}()
	return i
}

// go run -race problem.go
