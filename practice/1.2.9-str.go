package main

import (
	"fmt"
	"regexp"
)

var cnRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

func main() {
	str := "我爱go"
	StrFilterGetChinese(&str)
	fmt.Println(str)
}

func StrFilterGetChinese(src *string) {
	strNew := ""
	fmt.Println(*src)
	for _, c := range *src {
		if cnRegexp.MatchString(string(c)) {
			fmt.Println(string(c))
			fmt.Println(c)
			strNew += string(c)
		}
	}
	*src = strNew
}
