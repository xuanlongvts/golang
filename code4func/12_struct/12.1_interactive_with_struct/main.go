package main

import "fmt"

type StudentInfo struct {
	class string
	email string
	age int
}

type StudentMain struct {
	id int
	name string
	info StudentInfo
}

func main()  {

	// Zero value
	/*
		Golang se khoi tao gia tri mac dinh cua kieu du lieu
	*/
	var studenZero StudentMain
	fmt.Println("Zero value struct = ", studenZero)

	// anonymous fields
	type NoName struct{
		string
		int
	}
	var varNoName = NoName{"ABC", 90}
	fmt.Println("NoName = ", varNoName)

	// Nested struct
	student := StudentMain{
		id:   123,
		name: "LongLe",
		info: StudentInfo{
			class: "Lop chon",
			email: "abc@gmail.com",
			age: 90,
		},
	}
	fmt.Println("student nested = ", student)

	studentOne := StudentMain{
		id:   456,
		name: "HCM",
	}
	fmt.Println("student one = ", studentOne)

	// compare two struct
	/*
		1. struct chi so sanh duoc voi nhau khi chua kieu du lieu so sanh duoc, vd: int, string
		2. struct khong so sanh duoc voi nhau khi chua kieu du lieu khong so sanh duoc, vd: map
	*/
	type structOne struct{
		id int
		name string
	}
	st1 := structOne{
		id:   123,
		name: "abc",
	}
	st2 := structOne{
		id:   123,
		name: "abc",
	}
	if st1 == st2 {
		fmt.Println("st1 = st2")
	} else {
		fmt.Println("st1 !!!=== st2")
	}

	fmt.Println("=======================")
	//type structTwo struct {
	//	id int
	//	name string
	//	info map[int] int
	//}
	//st3 := structTwo{
	//	id: 1,
	//	name: "name",
	//	info: map[int] int {
	//		1: 1,
	//	},
	//}
	//st4 := structTwo{
	//	id:   1,
	//	name: "name",
	//	info: map[int]int{
	//		1: 1,
	//	},
	//}
	//if st3 == st4 {
	//	fmt.Println("st3 = st4")
	//} else {
	//	fmt.Println("st3 !!!=== st4")
	//}
}