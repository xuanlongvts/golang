package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 100

	fmt.Println(<-c)
}
