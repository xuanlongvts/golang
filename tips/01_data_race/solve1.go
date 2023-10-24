package main

import (
	"fmt"
	"sync"
)

// Way 1. Use WaitGroup

func main() {
	var i int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		i++
		wg.Done()
	}()
	wg.Wait()

	fmt.Println(i)
}
