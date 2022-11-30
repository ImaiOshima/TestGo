package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	s := "我是重负偶人aaa"
	var out []rune
	for _, r := range s {
		out = append(out, r)
	}
	fmt.Println(string(out))
}
