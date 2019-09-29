package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond * 100)
	// ctx, _ := context.WithTimeout(context.Background(), time.Second * 100)
	req, _ := http.NewRequest(http.MethodGet, "https://google.com/", nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed", err)
		return
	}
	fmt.Println("Request success, status:", res.StatusCode)
}