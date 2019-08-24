package main

import "fmt"

func main(){
	if number := 100; number == 100 {
		fmt.Println("number = ", number)
	}

	var number = 500
	switch number {
	case 1, 3, 5, 7:
		fmt.Println("so = 1")
	case 2:
		fmt.Println("so = 2")
	default:
		fmt.Println("so = ", number)
	}

	fmt.Println("===================")
	switch {
	case number > 400:
		fmt.Println("number > 400, number is ", number)
	case number < 400:
		fmt.Println("number < 400 ")
	}
}