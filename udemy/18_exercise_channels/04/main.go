package main

import "fmt"

func main() {
	c := make(chan int)

	r := gen(c)

	for i := range r {
		fmt.Println("i =", i)
	}

	fmt.Println("About exit")
}

func receive(c chan int) {
	for i := range c {
		fmt.Println("i =", i)
	}
}

func gen(q chan int) <-chan int {
	go func() {
		for i := 0; i < 10; i++ {
			q <- i
		}
		close(q)
	}()

	return q
}
