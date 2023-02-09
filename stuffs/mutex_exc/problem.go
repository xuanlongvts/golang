package main

import (
	"fmt"
	"math/rand"
	"time"
)

var count = 0

func process(n int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		temp := count
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
	}
	fmt.Println("Count after i=", n, " Count: ", count)
}

func main() {
	for i := 0; i < 3; i++ {
		go process(i)
	}
	fmt.Println("test log")
	time.Sleep(20 * time.Second)
	fmt.Println("Final count: ", count)
}
