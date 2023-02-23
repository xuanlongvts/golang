package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

var counter = 0

func process1(n int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		mu.Lock()

		temp := counter
		temp++
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		counter = temp

		mu.Unlock()
	}
	fmt.Println("Count after i=", n, " Count: ", counter)
}

func main() {
	for i := 0; i < 3; i++ {
		go process1(i)
	}
	fmt.Println("test log")
	time.Sleep(20 * time.Second)
	fmt.Println("Final count: ", counter)
}
