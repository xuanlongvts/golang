package main

import (
	"fmt"
	"encoding/xml"
)

func main()  {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name string `xml:"FullName"`
		Phone string
		Email []Email
		Groups []string `xml:"Groups>Value"`
		Address
	}

	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Nguyễn Thị Minh Không Khai</FullName>
			<Company>VinID</Company>
			<Email where="home">
				<Addr>abc@gmail.com</Addr>
			</Email>
			<Email where="work">
				<Addr>def@gmail.com</Addr>
			</Email>
			<Groups>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Groups>
			<City>HCM</City>
			<State>Q3</State>
		</Person>
	`

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
}