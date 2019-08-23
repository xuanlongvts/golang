package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("Exist")
}

// only send
func foo(c chan<- int) {
	c <- 40
}

// only recieve
func bar(c <-chan int) {
	fmt.Println("c: ", <-c)
}
