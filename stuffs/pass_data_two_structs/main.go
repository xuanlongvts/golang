package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Staff struct {
	User
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func main() {
	user1 := User{
		ID:   "1",
		Name: "Minh",
	}
	byteUser, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}

	var staff1 Staff
	err = json.Unmarshal(byteUser, &staff1)
	if err != nil {
		panic(err)
	}
	staff1.Address = "555 3/2 Quan 10"
	staff1.Phone = "2934728374"
	fmt.Println("staff1: ", staff1)

	fmt.Println("Way 2 ------------------------")
	var staff2 Staff
	err = PopulateDataTwoStructs(user1, &staff2)
	if err != nil {
		panic(err)
	}
	staff2.Address = "300 Le Hong Phong"
	staff2.Phone = "234234324"
	fmt.Println("staff2: ", staff2)
}

func PopulateDataTwoStructs[T any](original any, des *T) error {
	byteData, err := json.Marshal(original)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &des)
	if err != nil {
		return err
	}
	return nil
}
