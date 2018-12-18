package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type truct struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truct{
		vehicle: vehicle{
			door:  2,
			color: "white",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			door:  4,
			color: "Black",
		},
		luxury: false,
	}
	fmt.Println("t = ", t)
	fmt.Println("s = ", s)
	fmt.Println("door truct = ", t.door)
	fmt.Println("door sedan = ", s.door)
}
