package main

import "fmt"

type Bike struct {
	sound string
}

type Car struct {
	sound string
}

func (b Bike) SoundHorn() {
	fmt.Println(b.sound)
}

func (c Car) SoundHorn() {
	fmt.Println(c.sound)
}

type HornSounder interface {
	SoundHorn()
}

func main() {
	b := Bike{
		sound: "Ren ren",
	}
	b.SoundHorn()

	vehi := []HornSounder{
		Bike{"Ring"},
		Car{"BEEP"},
	}
	for _, v := range vehi {
		v.SoundHorn()
	}
}
