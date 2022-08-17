package main

import (
	"fmt"
)

func main() {
	//a := struct {
	//	x int
	//}{}
	//
	//a.x = 100
	//p := &a
	//p.x += 100
	//fmt.Println(p.x)

	type data struct {
		x int
		s string
	}
	var a data = data{1, "abc"}
	b := data{
		1,
		"abb",
	}
	c := []int{
		1, 2,
	}
	d := []int{
		1, 2, 3, 4, 5}
	fmt.Println(a, b, c, d)
}
