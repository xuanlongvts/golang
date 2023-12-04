package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr = [3]int{1, 2, 3}
	checkType(arr)

	var sli = []int{4, 5, 6}
	checkType(sli)
}

func checkType(t interface{}) {
	fmt.Println("---------------------")

	typ := reflect.TypeOf(t)
	fmt.Println("typ: ", typ.Kind())
	if typ.Kind() == reflect.Slice {
		for k, v := range t.([]int) {
			fmt.Println("slice, ", k, v)
		}
	}

	if typ.Kind() == reflect.Array {
		for k, v := range t.([3]int) {
			fmt.Println("arr, ", k, v)
		}
	}
}
