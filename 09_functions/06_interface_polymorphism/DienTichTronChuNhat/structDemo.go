package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func rectangleArea(r Rectangle) float64 {
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)

	return l * w
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 3}
	r := Rectangle{0, 0, 10, 10}
	/* or
	c := Circle{x: 0, y : 0, r : 5}
	r := Rectangle{x1 : 0, y1 : 10, x2 : 0, y2 : 10}
	*/
	fmt.Println(circleArea(c))
	fmt.Println(rectangleArea(r))
}
