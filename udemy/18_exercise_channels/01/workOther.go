package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 100

	fmt.Println("c: ", <-c)
}
