package  main

import "fmt"

func main(){

	// infinite loop

	/*
	for {
		fmt.Println("loop infinite")
	}
	*/

	for i, k := 1, 2; i < 10 && k < 10; i,k = i + 1, k + 1 {
		fmt.Println("i = ", i)
		fmt.Println("k = ", k)
	}
}