package main

import (
	"fmt"
	"strings"
)

type FilesInlineBody struct {
	Link string
	ID   string
}

var prefixIdImgs = map[string]string{
	"default": "",
	"gmail":   "cid:",
}

func combinePrefixToIdImgs(typeMail string, links []FilesInlineBody) []FilesInlineBody {
	data := []FilesInlineBody{}
	for k, v := range prefixIdImgs {
		for _, val := range links {
			if k == typeMail {
				newId := v + val.ID
				eachLink := FilesInlineBody{
					Link: val.Link,
					ID:   newId,
				}
				data = append(data, eachLink)
			}
		}
	}
	return data
}

func main() {
	var links = []FilesInlineBody{{
		Link: "http://localhost:1323/files_from_email/download.png",
		ID:   "ii_lhsx7ku01",
	}, {
		Link: "http://localhost:1323/files_from_email/image2.png",
		ID:   "ii_lhsx7ktu0",
	}}
	res := combinePrefixToIdImgs("gmail", links)
	//fmt.Println("res: ", res)

	var str = "<div dir=\"ltr\"><div><img src=\"cid:ii_lhsx7ktu0\" alt=\"image2.png\" width=\"250\" height=\"542\"><img src=\"cid:ii_lhsx7ku01\" alt=\"download.png\" width=\"542\" height=\"542\"><br></div><div><br></div><div>noi dung</div><br clear=\"all\"><div><div dir=\"ltr\" class=\"gmail_signature\" data-smartmail=\"gmail_signature\"><div dir=\"ltr\"><div><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\">Best regards,</span><br style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255)\"><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\">______________________________</span><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\">_</span><br style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255)\"></div><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\">Xuan Long</span><br style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255)\"><div><br style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255)\"><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\">Mobile: 096.928.2984</span><br style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255)\"><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\">Email:<span></span></span><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\"><span></span></span><span style=\"color:rgb(136,136,136);font-family:arial,sans-serif;font-size:13px;font-style:normal;font-variant:normal;font-weight:normal;letter-spacing:normal;line-height:normal;text-align:start;text-indent:0px;text-transform:none;white-space:normal;word-spacing:0px;background-color:rgb(255,255,255);display:inline!important;float:none\"><span> </span></span><a href=\"mailto:xuanlongvts@gmail.com\" target=\"_blank\">xuanlongvts@gmail.com</a></div></div></div></div></div>"

	var result = ""
	for _, v := range res {
		if result != "" {
			str = result
		}
		result = strings.Replace(str, v.ID, v.Link, 1)
	}
	fmt.Println("result: ", result)
}
