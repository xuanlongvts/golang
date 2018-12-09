package main

import "fmt"

var x int = 42
var y string = "Nong No"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println("s = ", s)
}

/*
Using the code from the previous exercise,
1. at the package level scope, assign the following values to the three variables a. for x assign 42 b. for y assign "James Bond" c. for z assign true
2. in func main a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of type string using the short declaration operator to a VARIABLE with the IDENTIFIER s
*/
