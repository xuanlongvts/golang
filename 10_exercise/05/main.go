package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) area() float64 {
	return s.length * s.length
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s.area())
}

func main() {
	circ := Circle{
		radius: 15.2,
	}
	squr := Square{
		length: 10,
	}
	info(circ)
	info(squr)
}
