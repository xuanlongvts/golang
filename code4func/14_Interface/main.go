package main

import "fmt"

// interface
// multiple interface
// Embed interface
// Empty interface

type Animal interface {
	speak()
}
type Movement interface {
	move()
}

// Embed interface
type NextAnimal interface {
	Animal
	Movement
}

type Dog struct {}

func (d Dog) speak() {
	fmt.Println("Go Go Go")
}

func (d Dog) move() {
	fmt.Println("Dog runs by four legs")
}

// empty
type data struct {
	index int
}
func goOut(i interface{}) {
	fmt.Println("Empty interface = ", i)
}

func main()  {
	var animal Animal
	animal = Dog{}
	animal.speak()

	fmt.Println("====================================")
	dog := Dog{}
	var mov Movement = dog
	mov.move()

	var ani Animal = dog
	ani.speak()

	fmt.Println("====================================")
	fmt.Println("Next Animal (Embed)")
	var nxtAnimal NextAnimal = dog
	nxtAnimal.speak()
	nxtAnimal.move()

	fmt.Println("====================================")
	goOut(10)
	goOut(1.2330)
	var dataStruct = data{index: 10}
	goOut(dataStruct)
}