package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := User{
		First: "Long",
		Last:  "Le",
		Age:   30,
		Sayings: []string{
			"Noi mot loi",
			"Noi loi hai",
		},
	}
	u2 := User{
		First: "Nong",
		Last:  "No",
		Age:   20,
		Sayings: []string{
			"Tao Lao Me Lao",
			"Xam Xi Du",
		},
	}
	listUser := []User{u1, u2}
	fmt.Println(listUser)

	err := json.NewEncoder(os.Stdout).Encode(listUser)
	if err != nil {
		fmt.Println("Something is wrong", err)
	}

}
