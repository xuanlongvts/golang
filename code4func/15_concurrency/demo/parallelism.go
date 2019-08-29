package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func createPizza(pizza int) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Create Pizza %d \n", pizza)
}

func timeTract(start time.Time, functionName string) {
	elapsed := time.Since(start)
	fmt.Println(functionName, " took ", elapsed)
}

func main()  {
	defer timeTract(time.Now(), "Build Pizzas")

	number_core_cpu := 3
	runtime.GOMAXPROCS(number_core_cpu)
	wg.Add(number_core_cpu)
	for i :=0; i < number_core_cpu; i++ {
		go createPizza(i)
	}
	wg.Wait()
}