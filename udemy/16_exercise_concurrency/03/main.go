package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	gs := 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Curr counter:", counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End counter:", counter)
}

/*
	There are two ways for test

	1. go run -race twoContext.go
	2. go build -race
		2.1 run ./03     to test
*/
