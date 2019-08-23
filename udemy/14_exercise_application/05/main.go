package main

import (
	"fmt"
	"sort"
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
	u3 := User{
		First: "BB",
		Last:  "CC",
		Age:   60,
		Sayings: []string{
			"Gi zay choi",
			"An no",
		},
	}
	listUser := []User{u1, u2, u3}
	fmt.Println(listUser)
	fmt.Println("====================================")

	sort.Slice(listUser, func(i, j int) bool {
		return listUser[i].Age < listUser[j].Age
	})
	for _, val := range listUser {
		fmt.Println(val.First, val.Last, ":", val.Age)
		for _, v := range val.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("====================================")
	sort.Slice(listUser, func(i, j int) bool {
		return listUser[i].First < listUser[j].First
	})
	for _, val := range listUser {
		fmt.Println(val.First, val.Last, ":", val.Age)
		sort.Strings(val.Sayings) // sort Sayings
		for _, v := range val.Sayings {
			fmt.Println("\t", v)
		}
	}
}
