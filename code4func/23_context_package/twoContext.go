package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 1)
	fmt.Println("Logic in main func")
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println("Logic in doSomething func")
	time.Sleep(time.Second * 3) // block this point
	fmt.Println("End after 3s")
}

// Yeu cau la sau 1s se cancel toan bo logic trong func doSomething
// this is wrong program