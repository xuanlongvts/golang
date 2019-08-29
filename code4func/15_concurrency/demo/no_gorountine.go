package main

import (
	"fmt"
	"time"
)

func createPizzaOne(pizza int) {
	time.Sleep(time.Second)
	fmt.Printf("Create a Pizza %d \n", pizza)
}

func timeTrackOne(start time.Time, functionName string)  {
	elapsed := time.Since(start)
	fmt.Println(functionName, " took ", elapsed)
}

func main()  {
	defer timeTrackOne(time.Now(), "Build Pizzas ")

	for i := 0; i < 3; i++ {
		createPizzaOne(i)
	}
}