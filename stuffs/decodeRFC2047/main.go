package main

import (
	"fmt"
	"github.com/jhillyerd/enmime"
	"mime"
	"strings"
)

func main() {
	//encodedFilename := "=?UTF-8?B?SFNLLVAtMDIzX1BMMDUgSMaw4bubbmcgZOG6q24ga2hhaSBiw6FvIHdvcmtmbG93ICgxKS5wZGY=?="
	//encodedFilename := "=?UTF-8?B?SFNLLVAtMDIzX1BMMDUgSMaw4bubbmcgZOG6q24ga2hhaSBiw6FvLnBkZg==?="
	encodedFilename := "=?UTF-8?B?SFNLLUhELTAzMiAtIEjGsOG7m25nIGThuqtuIHPhu60gZOG7pW5nIERhaWx5IFRhc2sgdOG6oWkgSA==?= =?UTF-8?B?YXNha2kgV29yayAtIFZlci4gMDAucGRm?="

	decoder := new(mime.WordDecoder)
	decodedFilename, err := decoder.Decode(encodedFilename)
	if err != nil {
		fmt.Println("Error decoding filename:", err)
	}
	fmt.Println("mime Filename:", decodedFilename)

	decodedFilename = enmime.DecodeRFC2047(encodedFilename)
	fmt.Println("enmime Filename:", decodedFilename)

	fmt.Println("---------------------------")
	var str = "name:  attachment; filename=\"=?UTF-8?B?SFNLLVAtMDIzX1BMMDUgSMaw4bubbmcgZOG6q24ga2hhaSBiw6FvLnBkZg==?=\""
	res1 := strings.Split(str, "\"")
	fmt.Println("res1: ", res1[1])
}
