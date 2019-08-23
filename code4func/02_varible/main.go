package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println("Number = ", number)

	var numberOne int = 20
	fmt.Println("Number = ", numberOne)

	var stringShow= "longle"
	fmt.Println("String show =", stringShow)

	stringShowTwo := "NongNo"
	fmt.Println("Strig show 2 = ", stringShowTwo)

	// Cach 1: Khai bao nhieu bien voi cung kieu du lieu
	var a, b int
	a = 10
	b = 20
	fmt.Println("a = ", a, ":; b = ", b)

	var c, d int = 30, 40
	fmt.Println("c = ", c, ":; d = ", d)

	var e, f = 50, 60
	fmt.Println("e = ", e, ":; f = ", f)



	// Cach 2: Khai bao nhieu bien voi khac kieu du lieu
	var (
		name string
		address string
		age int
	)
	name = "abc"
	address = "nguyen thi minh khai"
	age = 20
	fmt.Println("name = ", name, ", address = ", address, ", age = ", age)

	var (
		name1 string = "name1"
		address1 string = "address1"
		age1 int = 100
	)
	fmt.Println("name1 = ", name1, ", address1 = ", address1, ", age1 = ", age1)

	var name2, address2, age2 = "name2", "address2", 50
	fmt.Println("name2 = ", name2, ", address2 = ", address2, ", age2 = ", age2)
}