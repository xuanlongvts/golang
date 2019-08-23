package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p := person{
		first:      "Long",
		last:       "Le",
		favFlavors: []string{"football", "music", "code"},
	}
	fmt.Println("person=", p)
	fmt.Println("first ", p.first)
	fmt.Println("last ", p.last)
	fmt.Println("==============================")
	for i, v := range p.favFlavors {
		fmt.Println("i, v = ", i, v)
	}
}
