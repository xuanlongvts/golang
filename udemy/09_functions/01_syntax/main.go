package main

import "fmt"

func main() {
	foo()
	bar("Long Le")

	s1 := woo("The world")
	fmt.Println("s1=", s1)

	fmt.Println("====================")
	x, y := person("Xuan", "Long")
	fmt.Print(x, ": ")
	fmt.Println(y)

}

func foo() {
	fmt.Println("From foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func person(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, ", says Hello")
	b := false
	return a, b
}
