package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goruntines:", runtime.NumGoroutine())

	counter := 0
	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)

	var mut sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mut.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mut.Unlock()
			wg.Done()
			fmt.Println("Goruntines11:", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Goruntines 22:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
