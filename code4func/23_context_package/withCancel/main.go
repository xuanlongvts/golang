package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(time.Second * 2, func() {
		fmt.Println("Here in 2s will cancel")
		cancel()
	})

	select {
		case <- ctx.Done():
			fmt.Println("Done")
	}
}