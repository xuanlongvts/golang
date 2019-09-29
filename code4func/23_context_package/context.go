package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)

	// Cach 1
	//defer cancel()

	// Cach 2
	time.AfterFunc(time.Second, func() {
		fmt.Println("Cancel before 3s")
		cancel()
	})

	select {
		case <- ctx.Done():
			fmt.Println("Err done = ", ctx.Err())
	}
}