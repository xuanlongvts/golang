package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 1)
	fmt.Println("Logic in main func")

	doSomethingMore(ctx)
}

func doSomethingMore(ctx context.Context) {
	canceledChannel := make(chan bool)
	go func() {
		select {
		case <- ctx.Done():
			fmt.Println("Error = ", ctx.Err())
			canceledChannel <- true
			return
		}
	}()

	isCanceledChannel := <- canceledChannel
	if isCanceledChannel {
		close(canceledChannel)
		return
	}

	time.Sleep(time.Second * 3)
	fmt.Println("Logic end after 3s in doSomething func")
}

// Yeu cau la sau 1s se cancel toan bo logic trong func doSomething