package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxParent, cancel :=  context.WithCancel(context.Background())
	ctxTimeout, _ := context.WithTimeout(ctxParent, time.Second * 5) // context after 5s will break

	//defer cancel()
	time.AfterFunc(time.Second, func() {
		fmt.Println("After 1s will break")
		cancel()
	})

	select {
		case <- ctxTimeout.Done():
			fmt.Println("Timeout")
	}
}