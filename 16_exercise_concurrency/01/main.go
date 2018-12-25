package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin CPUs: ", runtime.NumCPU())
	fmt.Println("begin gs: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("One")
		wg.Done()
	}()

	go func() {
		fmt.Println("Two")
		wg.Done()
	}()
	fmt.Println("Middle CPUs: ", runtime.NumCPU())
	fmt.Println("Middle gs: ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Exist")
	fmt.Println("End CPUs: ", runtime.NumCPU())
	fmt.Println("End gs: ", runtime.NumGoroutine())
}
