package main

import "fmt"

var x = 30

// Static programing language, not dynamic programing language like javascript
var y = "Long Lờ" // var y string = "Long Lờ"

var strA = `Long
Le

Xuan
`

func main() {
	fmt.Println("x = ", x)
	fmt.Printf("%T", x)

	fmt.Println("y = ", y)
	fmt.Printf("%T", y)

	// y = 5 error

	fmt.Println("y = ", strA)
	fmt.Printf("%T", strA)
}
