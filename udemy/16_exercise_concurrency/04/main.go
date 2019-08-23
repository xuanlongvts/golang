package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var counter = 0
	var gs = 10
	wg.Add(gs)

	var muT sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			muT.Lock()
			v := counter
			v++
			counter = v
			// muT.Unlock()
			fmt.Println("Curr counter:", counter)
			muT.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("========================")
	fmt.Println("End counter:", counter)
}
