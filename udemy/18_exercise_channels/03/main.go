package main

import "fmt"

func main() {
	c := gen()

	receive(c)

	fmt.Println("About exist")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("c =", v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}
