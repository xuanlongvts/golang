package main

import "fmt"

type Speak interface {
	speak()
}
type Movement interface {
	move()
}
type Color interface {
	color()
}

// Embed interface
type Animal interface {
	Speak
	Movement
	Color
}

type Bird struct {}

func (bir *Bird) speak() {
	fmt.Println("Rit rit rit")
}

func (bir *Bird) move() {
	fmt.Println("Bay hang dam")
}

func (bir *Bird) color() {
	fmt.Println("Mau den")
}

func main() {
	var bird Animal = &Bird{}
	bird.speak()
	bird.move()
	bird.color()
}