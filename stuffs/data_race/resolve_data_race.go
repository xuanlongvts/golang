package main

import (
	"fmt"
	"sync"
)

func main() {
	var i = 2
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		mu.Lock()
		i = i + 2
		mu.Unlock()

		wg.Done()
	}()
	go func() {
		mu.Lock()
		i = i * 3
		mu.Unlock()

		wg.Done()
	}()
	wg.Wait()
	fmt.Println("result: ", i)
}

// go run resolve_data_race.go
// this way resolve data race but still happen race condition
