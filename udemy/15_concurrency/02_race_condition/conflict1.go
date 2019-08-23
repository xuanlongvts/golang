package main

import (
	"fmt"
	"sync"
)

var (
	counter int64
	mutex   sync.Mutex
)

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			for ii := 0; ii < 100; ii++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	var str string
	fmt.Scanln(&str)
	fmt.Println("counter =", counter)
}
