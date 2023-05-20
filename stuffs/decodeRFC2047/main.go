package main

import (
	"fmt"
	"strings"
)

func main() {
	//encodedFilename := "=?UTF-8?B?SFNLLVAtMDIzX1BMMDUgSMaw4bubbmcgZOG6q24ga2hhaSBiw6FvIHdvcmtmbG93ICgxKS5wZGY=?="
	//encodedFilename := "=?UTF-8?B?SFNLLVAtMDIzX1BMMDUgSMaw4bubbmcgZOG6q24ga2hhaSBiw6FvLnBkZg==?="
	//
	//decodedFilename := enmime.DecodeRFC2047(encodedFilename)
	//
	//fmt.Println("Decoded Filename:", decodedFilename)

	var str = "name:  attachment; filename=\"=?UTF-8?B?SFNLLVAtMDIzX1BMMDUgSMaw4bubbmcgZOG6q24ga2hhaSBiw6FvLnBkZg==?=\""
	res1 := strings.Split(str, "\"")

	fmt.Println("res1: ", res1[1])
}
