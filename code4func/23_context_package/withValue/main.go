package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "number", 10)

	A(ctx)
}

func A(ctx context.Context) {
	if value := ctx.Value("number"); value != nil {
		ctxA := context.WithValue(ctx,"number1", 5)
		B(ctxA)
	}
}

func B(ctx context.Context)  {
	valOne := ctx.Value("number").(int)
	valTwo := ctx.Value("number1").(int)

	fmt.Println("Context value total = ", valOne + valTwo)
}