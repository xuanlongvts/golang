package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 40
	c <- 100
	fmt.Println("c1:", <-c)
	fmt.Println("c2:", <-c)
	fmt.Println("==========")
	fmt.Printf("%T\n", c)
}
