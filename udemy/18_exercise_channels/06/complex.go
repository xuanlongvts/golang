package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 4
	y := 4

	c := gen(x, y)

	for index := 0; index < x*y; index++ {
		fmt.Println("index=", index, ",\t value c:", <-c)
	}

	fmt.Println("End goroutines: ", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for ii := 0; ii < y; ii++ {
				c <- ii
			}
		}()

		fmt.Println("Number goruntines:", runtime.NumGoroutine())
	}

	return c
}
