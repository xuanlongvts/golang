package main

import (
	"fmt"
	"sync/atomic"
)

var counter int64

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			for ii := 0; ii < 100; ii++ {
				atomic.AddInt64(&counter, 1)
				// counter++
			}
		}()
	}

	var str string
	fmt.Scanln(&str)
	fmt.Println("counter =", counter)
}
