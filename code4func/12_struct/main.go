package main

import "fmt"

type Student struct {
	id int
	name string
}

func main() {
	// Cach1: named
	var stdOne = Student{
		id:   123,
		name: "LongLe",
	}
	fmt.Println("Student 1 = ", stdOne)
	fmt.Println("id 1: ", stdOne.id)
	fmt.Println("name 1: ", stdOne.name)

	// Cach2: short
	var stdTwo = Student{456, "NongNo"}
	fmt.Println("Student 2 = ", stdTwo)

	// Cach3: declation and init
	var stdThree = struct {
		age int
		add string
	}{
		age: 50,
		add: "HCM",
	}
	fmt.Println("Student 3 = ", stdThree)

	// anonymous struct
	var anonymous = struct {
		email string
		age int
	}{
		"abc@gmailcom",
		20,
	}
	fmt.Println("anonymous = ", anonymous)


	// pointer tro toi struct
	fmt.Println("================")
	var pointerOne = &Student{
		id: 999,
		name: "ABCDE",
	}
	fmt.Println("pointerOne = ", pointerOne, ", hexa address = ", &pointerOne)
	fmt.Println("pointerOne id = ", pointerOne.id, ", or id = ", (*pointerOne).id)
	fmt.Println("pointerOne name = ", pointerOne.name, ", or name = ", (*pointerOne).name)
}