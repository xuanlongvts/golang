package main

import "fmt"

// Khong duoc goi mot gia tri vao channel da close

func main() {
	queueOne := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			if i > 5 {
				close(queueOne)
				break
			}
			queueOne <- i
		}
		//close(queueOne)
	}()

	for value := range queueOne {
		fmt.Println("v1 = ", value)
	}
}