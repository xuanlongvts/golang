package main

import (
	"fmt"
	"time"
)

func main() {
	var i = 2

	go func() {
		i = i * 3
	}()
	go func() {
		i = i + 2
	}()

	time.Sleep(time.Second)
	fmt.Println("result: : ", i)
}

// go run -race problem.go
