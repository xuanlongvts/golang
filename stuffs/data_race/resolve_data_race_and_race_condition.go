package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var i int = 2
	ch := make(chan int)

	go func() {
		wg.Add(1)
		go func() {
			i = i + 2
			ch <- i

			wg.Done()
		}()

		wg.Wait()
		wg.Add(1)
		go func() {
			i = i * 3
			ch <- i

			wg.Done()
		}()
		wg.Wait()

		close(ch)
	}()

	for v := range ch {
		res := v
		fmt.Println("res: ", res)
	}
}
