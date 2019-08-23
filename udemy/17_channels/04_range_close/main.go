package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println("chan: ", v)
	}
}

func foo(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
