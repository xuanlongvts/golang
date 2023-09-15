package main

import (
	"fmt"
	"regexp"
	"strings"
)

func parseTag(strData string) (listLinks []string, listHashtag []string) {
	var regexLink = `(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`
	var regexTagText = `(#+[a-zA-Z0-9(_)]{1,})`

	reLink := regexp.MustCompile(regexLink)
	listLinResks := reLink.FindAllString(strData, -1)

	for _, v := range listLinResks {
		aaa := strings.Replace(strData, v, "", -1)
		fmt.Println("aaa: ", aaa)
	}

	reTagText := regexp.MustCompile(regexTagText)
	reTagTextRes := reTagText.FindAllString(strData, -1)

	return listLinResks, reTagTextRes
}

func main() {
	strData := " #fff #sdf https://php-play.dev/?c=DwfgDgFmBQAkBOBTAzgVw#DYBcAEBebYAhvMo#gCqEDm https://genk.vn/sdfds"
	listLinks, listHashtag := parseTag(strData)
	fmt.Println("listLinks: ", listLinks)
	fmt.Println("listHashtag: ", listHashtag)
}
