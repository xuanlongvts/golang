package main

import "fmt"

var x = 42

func main() {
	fmt.Println("x = ", x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)

	x = 911
	fmt.Printf("%#x\n", x)
	fmt.Printf("%#x\t%b\t%x\n", x, x, x)

	z := fmt.Sprintf("%#x\t%b\t%x\n", x, x, x)
	fmt.Println(z)
	fmt.Printf("%v", x)
}
