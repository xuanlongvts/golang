package main

import "fmt"

func main() {
	c := make(<-chan int)
	// solution: c := make(chan int)

	go func() {
		c <- 1000
	}()

	fmt.Println("c =", <-c)
}
