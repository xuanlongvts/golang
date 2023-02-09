package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produceRandomInt(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func writeFile(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println("Create err: ", err)
		return
	}

	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println("Fprintln err: ", err)
			f.Close()
			done <- false
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println("Close err: ", err)
		done <- false
		return
	}
	done <- true
}

func main() {
	var data = make(chan int)
	var done = make(chan bool)
	wg := sync.WaitGroup{}
	fmt.Println("---> 111 ")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produceRandomInt(data, &wg)
	}
	fmt.Println("---> 222 ")
	go writeFile(data, done)
	fmt.Println("---> 333 ")
	go func() {
		wg.Wait()
		close(data)
	}()
	fmt.Println("---> 444 ")
	d := <-done
	if d == true {
		fmt.Println("Write file succeed")
	} else {
		fmt.Println("Write file failed")
	}
}
