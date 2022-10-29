package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func init() {

}

var (
	ServiceName = "golang-redis-docker"
	port        = "8100"
)

var ctx = context.Background()

func main() {
	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(fmt.Sprintf("ping ok %s", ServiceName)))
	// })
	// fmt.Println("start service with port: ", port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err1 := rdb.Set(ctx, "fullname1", "longle", time.Duration(time.Second*5)).Err()
	if err1 != nil {
		panic(err1)
	}

	err2 := rdb.Set(ctx, "fullname2", "Nguyen Van Tuan", time.Duration(time.Second*20)).Err()
	if err2 != nil {
		panic(err2)
	}

	val1, e1 := rdb.Get(ctx, "fullname1").Result()
	if e1 != nil {
		panic(e1)
	}
	fmt.Println("val1:", val1)

}
