package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // only recieve
	cs := make(chan<- int) // only send

	fmt.Printf("c: %T\n", c)
	fmt.Printf("cr: %T\n", cr)
	fmt.Printf("cs: %T\n", cs)

	fmt.Println("====================")
	// general to specific convert
	cr = c
	cs = c
	fmt.Printf("c: %T\n", (<-chan int)(c))
	fmt.Printf("c: %T\n", (chan<- int)(c))
}
