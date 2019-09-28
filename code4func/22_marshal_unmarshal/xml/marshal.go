package main

import (
	"fmt"
	"os"
	"encoding/xml"
)

func main() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName xml.Name `xml:"person"`
		Id int `xml:"id,attr"`
		FirstName string `xml:"name>first"`
		LastName string	`xml:"name>last"`
		Age int `xml:"age"`
		Height float32 `xml:"height,omitempty"`
		Married bool
		Address
		Comment string `xml:"comment"`
	}

	v := &Person{
		Id: 10,
		FirstName: "A",
		LastName: "Nguyễn Văn",
		Age: 30,
		Height: 1.7,
		Married: true,
	}
	v.Comment = "Anh nay giau qua"
	v.Address = Address{
		City:  "HCM",
		State: "Gò Vấp",
	}

	//output, err := xml.MarshalIndent(v, " ", " ")
	//if err != nil {
	//	fmt.Printf("error: %v\n", err)
	//}
	//os.Stdout.Write(output)

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent(" ", " ")
	if err := enc.Encode(v); err != nil {
		fmt.Println("error: $v\n", err)
	}
}