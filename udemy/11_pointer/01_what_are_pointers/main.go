package main

import "fmt"

func main() {
	a := 40
	fmt.Println("value a = ", a)
	fmt.Println("address a = ", &a) // & give an address

	fmt.Println("==============")
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a // or  b := &a
	fmt.Println("value b = ", b)
	fmt.Println("value b = ", *b) // * give the value store at an address when we have an address

	fmt.Println("==============")
	fmt.Println("value a = ", *&a)

	*b = 4000
	fmt.Println("============== after assign for b ")
	fmt.Println("value a = ", a)
}
