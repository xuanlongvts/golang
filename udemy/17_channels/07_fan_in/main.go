package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring2("Nong"), boring2("No"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("About exist")
}

func boring2(mess string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", mess, i)

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
