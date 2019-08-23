package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"Name"`
	Order string `json:"Order"`
}

func main() {
	s := []byte(`[{"Name": "Platypus", "Order": "Monotremata"},{"Name": "Quoll", "Order": "Dasyuromorphia"}]`)
	fmt.Println("s =", s)

	var People []Person
	err := json.Unmarshal([]byte(s), &People)
	if err != nil {
		fmt.Println("Err =", err)
	}
	fmt.Println("People =", People)

	for _, v := range People {
		fmt.Println("name:", v.Name, "\t order: ", v.Order)
	}
}
