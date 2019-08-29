package main

import (
	"fmt"
	"runtime"
	"time"
)

func go1() {
	fmt.Println("Go routines 1")
}

func go2() {
	fmt.Println("Go routines 2")
}

func main() {
	go go1()

	// count quantity goroutines
	ng := runtime.NumGoroutine()
	fmt.Println("Quantity goroutines = ", ng)

	time.Sleep(time.Second)
}