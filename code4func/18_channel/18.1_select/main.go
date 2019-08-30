package main

import "fmt"

func main() {
	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}

		done <- true
	}()

	for {
		select {
			case v := <-queue:
				fmt.Println("v = ", v)
			case <-done:
				fmt.Println("Done")
				return
		}
	}
}