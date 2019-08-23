package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println("x = ", x)
	fmt.Println("length x = ", len(x))
	fmt.Println("capacity x = ", cap(x))

	x[0] = 40
	x[9] = 999
	fmt.Println("after x = ", x)
	fmt.Println("after length x = ", len(x))
	fmt.Println("after capacity x = ", cap(x))

	fmt.Println("========================================")
	x = append(x, 14)
	fmt.Println("x = ", x)
	fmt.Println("length x = ", len(x))
	fmt.Println("capacity x = ", cap(x))

	x = append(x, 15)
	x = append(x, 16)
	x = append(x, 17)

	fmt.Println("--------------------------")
	fmt.Println("x = ", x)
	fmt.Println("length x = ", len(x))
	fmt.Println("capacity x = ", cap(x)) // capacity auto double size
}
