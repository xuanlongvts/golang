package main

import (
	"fmt"
	"encoding/xml"
)

type Plant struct {
	XMLName xml.Name	`xml:"plant"`
	Id		int 		`xml:"id, attr"`
	Name	string		`xml:"name"`
	Origin	string		`xml:"origin"`
}