package main

import (
	"fmt"
	"github.com/leekchan/accounting"
	"go-module/c4f"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))

	c4f.Println("this is text for color")
}
