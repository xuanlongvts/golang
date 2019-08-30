package main

import "fmt"

/*
	1. Unbuffered channel
	2. Buffered channel
	3. select
	4. Close channel
*/

func main()  {
	// 1. Unbuffered channel
	ch := make(chan int)

	go func() {
		ch <- 100 // this is point block
		fmt.Println("Sent unbuffered channel")
	}()

	fmt.Println(<-ch) // this is point block
	fmt.Println("Done")

	fmt.Println("========================")
	// 2. Buffered channel
	// Note: Buffered channel is not block any line code
	chBuff := make(chan int, 2) // if we replace 2 to 0 is the with   Unbuffered channel
	chBuff <- 1
	chBuff <- 2
	close(chBuff)
	fmt.Println(<-chBuff)
	fmt.Println(<-chBuff)
	fmt.Println(<-chBuff)
}