package main

import (
	"fmt"
	"sync"
)

func go1()  {
	fmt.Println("Go routines 111")
	wg.Done()
}

func go2() {
	fmt.Println("Go routines 222")
	wg.Done()
}

var wg sync.WaitGroup
func main() {
	fmt.Println("====== Begin")

	wg.Add(2)
	go go1()
	go go2()
	wg.Wait()

	fmt.Println("====== End")
}