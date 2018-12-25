package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Long","Last":"Le","Age":30},{"First":"Nong","Last":"No","Age":20}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs) // byte
	fmt.Println("value bs =", string(bs))

	fmt.Println("===================================")
	var peo []Person // per := []Person{}
	err := json.Unmarshal(bs, &peo)
	if err != nil {
		fmt.Println("Error message:", err)
	}
	fmt.Println(" All data person =", peo)
	fmt.Println("===================================")
	for i, v := range peo {
		fmt.Println("Person ", i)
		fmt.Println("\t", v.First, v.Last, ", age", v.Age)
	}
}
