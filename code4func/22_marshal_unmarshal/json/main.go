package main

import (
	"fmt"
	"encoding/json"
)

type Employee struct {
	Name string
	Age int
	Salary int
}

type Response struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Salary int `json:"salary"`
}

func main() {
	fmt.Println("Marshal Json")

	// Creating object using Employee Struct.
	empObj := Employee{
		Name: "Nguyễn Văn A",
		Age: 30,
		Salary: 3000,
	}

	// Convert Struct Object into Byte data.
	emp, _ := json.Marshal(empObj)

	// Convert Byte data into json string to display data.
	fmt.Println("Marshal struct to data = " , string(emp))



	fmt.Println("==================================================")
	fmt.Println("Unmarshal Json")
	strEmp := `{"Name":"Rachel", "Age":24, "Salary":344444}`

	// Creating json string into byte code.
	bytes := []byte(strEmp)

	// Create empty Response struct and assign res variable.
	var res Response

	// Unmarshal by passing a pointer to an empty structs.
	json.Unmarshal(bytes, &res)

	// Print the Struct Name value.
	fmt.Println("Response: Name = ", res.Name)
}