package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (C *Circle) area() float64 {
	return math.Pi * C.radius * C.radius
}

type Shape interface {
	area() float64
}

func info(S Shape) {
	fmt.Println("Area ", S.area())
}

func main() {
	var C = Circle{5}
	info(&C)
}
