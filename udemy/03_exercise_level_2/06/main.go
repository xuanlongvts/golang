package main

import (
	"fmt"
)

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("c ", c)
	fmt.Println("d ", d)
}
