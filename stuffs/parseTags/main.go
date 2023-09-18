package main

import (
	"fmt"
	"regexp"
	"strings"
)

func parseTag(strData string) (listHashtag []string) {
	var regexLink = `(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`
	var regexTagText = `(#+[a-zA-Z0-9(_)]{1,})`

	reLink := regexp.MustCompile(regexLink)
	listLinResks := reLink.FindAllString(strData, -1)

	var newStrData string
	for _, v := range listLinResks {
		newStrData = strings.Replace(strData, v, "", -1)
		strData = newStrData
	}

	reTagText := regexp.MustCompile(regexTagText)
	reTagTextRes := reTagText.FindAllString(strData, -1)

	return reTagTextRes
}

func main() {
	strData := " #fff #sdf https://php-play.dev/?c=DwfgDgFmBQAkBOBTAzgVw#DYBcAEBebYAhvMo#gCqEDm https://genk.vn/sdfds"
	listHashtag := parseTag(strData)

	fmt.Println("--------------")
	fmt.Println("listHashtag: ", listHashtag)
}
