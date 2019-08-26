package  main

import "fmt"

// Variadic function
/*
	1. declation variadic function
	2. pass a slice to a variadic function
	3. change slice (Khi truyen vao variadic function mot slice thi Golang se thao tac truc tiep tren slice do luon, nen lam thay doi gia tri cua slice ben ngoai truyen vao)
*/

func addItem(item int, list ...int){
	// 100, 200, 300 -> []int {100, 200, 300}
	list = append(list, item)
	fmt.Println("list item = ", list)
}

func changeSlice(list ...int) {
	list[0] = 999
}

func main() {
	addItem(777, 100, 200, 300)

	var list = []int {1, 2, 3, 4}
	addItem(888, list...)

	changeSlice(list...)
	fmt.Println("List after change = ", list)
}