package main

import "fmt"

type Student struct{
	name string
}

// define method
// format: func (t Type) methodName(params) return { body code }
// (t Type) ---> Receiver, co  2 kieu Receiver
// 1. Value receiver
// 2. Pointer receiver

// Về bản chất method cũng là một function, nhưng khi nói đến method thì nó thuộc đối tượng nào

// 1. Value receiver
/*
	Khi su dung value receiver thi golang se copy cai struct truyen vao thanh mot struct moi nen co dia chi khac nhau, va su dung gia tri copy do de thao tac data,
	nen khong lam thay doi gia tri goc
*/
func (s Student) getName() string{
	return s.name
}

// 1. Value receiver
func (s Student) changeName() {
	fmt.Printf("Address of st2 = %p", &s)
	s.name = "Name change"
}

// 2. Pointer receiver
func (s *Student) changeNameTwo(){
	fmt.Printf("Address of st pointer = %p", s)
	s.name = "AAA BBB CCC DDD ====> Change"
}

// non-struct
type Chuoi string
func (s Chuoi) appenString(str string) string{
	return fmt.Sprintf("%s%s", str, s )
}

func main()  {
	student := Student{name: "LongLe"}
	name := student.getName()
	fmt.Println("name = ", name)

	// value receiver
	st1 := Student{name: "Ryan"}
	fmt.Printf("Address of st1 = %p", &st1)
	fmt.Println("")
	st1.changeName()
	fmt.Println("")
	fmt.Println("st1 name = ", st1.name)

	// pointer receiver
	fmt.Println("==================================")
	st2 := &Student{name: "LLLLLLLLLL"}
	fmt.Printf("Address of st pointer = %p", st2)
	fmt.Println("")
	st2.changeNameTwo()
	fmt.Println("")
	fmt.Println("st2 name = ", st2.name)

	// non-struct
	fmt.Println("==================================")
	st3 := Chuoi("Nongno")
	newStr := st3.appenString("No No ==> ")
	fmt.Println("newStr = ", newStr)
}