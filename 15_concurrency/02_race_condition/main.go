package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			fmt.Println("Goruntines11:", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Goruntines2222 :", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}
