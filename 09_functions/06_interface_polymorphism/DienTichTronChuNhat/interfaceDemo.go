package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
}

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)

	return l * w
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, v := range shapes {
		area += v.area()
	}

	return area
}

func main() {
	c := Circle{0, 0, 3}
	r := Rectangle{0, 0, 10, 10}

	areaC := totalArea(&c)
	areaR := totalArea(&r)
	fmt.Println("c=", areaC) // each area
	fmt.Println("r=", areaR) // each area

	totalT := totalArea(&c, &r)
	fmt.Println("total=", totalT) // total
}
