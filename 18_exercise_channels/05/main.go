package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 100
	}()

	v, ok := <-c
	fmt.Println("v=", v, ", ok=", ok)

	close(c)

	fmt.Println("==================")
	v, ok = <-c
	fmt.Println("v=", v, ", ok=", ok)
}
