package main

import "fmt"

func main() {

	func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i=", i)
		}
	}()

	fmt.Println("Done")
}
