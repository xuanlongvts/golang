package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 3
	y := 3
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for ii := 0; ii < y; ii++ {
				c <- ii

			}
		}()
		fmt.Println("Number gorountine:", runtime.NumGoroutine())
	}

	for k := 0; k < x*y; k++ {
		fmt.Println(k, ", \t", <-c)
	}

	fmt.Println("End number gorountine:", runtime.NumGoroutine())
}
