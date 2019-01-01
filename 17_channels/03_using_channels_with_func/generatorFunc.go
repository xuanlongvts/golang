package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nong := boring1("NongNo")
	joe := boring1("joe")
	ann := boring1("ann")

	for i := 0; i < 5; i++ {
		fmt.Println(<-nong)
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("About exist")
}

func boring1(mess string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", mess, i)

			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
	}()
	return c
}
