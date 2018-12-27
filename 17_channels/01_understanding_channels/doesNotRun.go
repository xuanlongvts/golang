package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 40
	fmt.Println(<-c)
}
