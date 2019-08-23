package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	go boring("Nong", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-c)
	}

	fmt.Println("You're boring, I'm leaving")

}

func boring(mess string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", mess, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
