package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 40
	}()

	fmt.Println("Channel C: ", <-c)
}
