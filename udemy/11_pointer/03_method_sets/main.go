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

func (C *Circle) area() float64 {
	return math.Pi * C.radius * C.radius
}

func (S *Square) area() float64 {
	return S.length * S.length
}

type Shape interface {
	area() float64
}

func infor(S Shape) {
	fmt.Println("area = ", S.area())
}

func main() {
	// cir1 := Circle{5}
	cir2 := Circle{7}
	infor(&cir2)

	sqr1 := Square{5}
	// sqr2 := Square{9}
	infor(&sqr1)
}
