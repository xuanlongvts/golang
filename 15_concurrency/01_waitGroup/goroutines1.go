package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Bat dau cac goroutines")
	go func() {
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
				time.Sleep(time.Millisecond * 10)
			}
		}
	}()

	go func() {
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
				time.Sleep(time.Millisecond * 10)
			}
		}
	}()

	fmt.Println("Doi ket thuc cac goroutines")
	var input string
	fmt.Scanln(&input)
	fmt.Println("End")
}
