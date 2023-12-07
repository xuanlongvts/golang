package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := `abc test testing edited @[Nguyễn Nhật Lam](user_id:1198)  @[Nguyễn Văn Lam](user_id:1199)`

	arr, _ := regexp.Compile(`(?im)@\[([^\]]+?)\]\(user_id:([^\]]+?)\)`)
	for _, match := range arr.FindAllStringSubmatch(str, -1) {
		fmt.Println("--------------------")
		fmt.Println(match[1], match[2])
	}
}
