package main

import "fmt"

// Way 2. Use Channel

func main() {
	var i int
	c := make(chan int)
	go increase(i, c)
	result := <-c
	fmt.Println(result)
}

func increase(i int, c chan int) {
	c <- i + 1
}
