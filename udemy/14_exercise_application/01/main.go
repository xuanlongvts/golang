package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// Note: We want to create a json, we have to declare type uppercase

func main() {
	u1 := User{
		Name: "Long",
		Age:  30,
	}
	u2 := User{
		Name: "Le",
		Age:  50,
	}
	listUser := []User{u1, u2}
	fmt.Println("listUser =", listUser)

	bs, err := json.Marshal(listUser) // Marshal like JSON.parse of javascript
	if err != nil {
		fmt.Println("Error happen ", err)
		return
	}

	fmt.Println("bs json =", string(bs))
}
