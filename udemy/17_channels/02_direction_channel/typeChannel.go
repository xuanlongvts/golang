package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // only recieve data
	cs := make(chan<- int) // only send data

	fmt.Printf("c: %T\n", c)
	fmt.Printf("cr: %T\n", cr)
	fmt.Printf("cs: %T\n", cs)
}
