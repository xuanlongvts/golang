package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	str := `abc test testing edited @[Nguyễn Nhật Lam](user_id:1198)  @[Nguyễn Văn Lam](user_id:1199)`

	arr, _ := regexp.Compile(`(?im)@\[([^\]]+?)\]\(user_id:([^\]]+?)\)`)
	for _, match := range arr.FindAllString(str, -1) {
		fmt.Println("--------------------")
		splStr := strings.Split(match, "@[")[1]
		getFullName := strings.Split(splStr, "](user_id:")[0]
		fmt.Println(getFullName)

		getID := strings.Split(splStr, "](user_id:")[1]
		getID = strings.Split(getID, ")")[0]
		fmt.Println(getID)
	}
}
