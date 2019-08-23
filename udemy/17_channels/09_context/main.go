package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context: ", ctx)
	fmt.Println("Context err: ", ctx.Err())
	fmt.Printf("Context type %T \n", ctx)
	fmt.Println("=================")

	// ctx, _ = context.WithCancel(ctx)
	// fmt.Println("Context: ", ctx)
	// fmt.Println("Context err: ", ctx.Err())
	// fmt.Printf("Context type %T \n", ctx)

	// fmt.Println("=================")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("Context cancel: ", cancel)
	fmt.Printf("Cancel type %T \n", cancel)

	fmt.Println("-----------------------")

	cancel()
	fmt.Println("Context: ", ctx)
	fmt.Println("Context err: ", ctx.Err())
	fmt.Printf("Context type %T \n", ctx) // different
	fmt.Println("Context cancel: ", cancel)
	fmt.Printf("Cancel type %T \n", cancel)
}
