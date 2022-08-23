package main

import "fmt"

func main() {
	type user struct {
		name string
		age  byte
	}

	d := [...]user{
		{"Tom", 20},
		{"Mary", 18},
	}

	fmt.Printf("%#v\n", d)
}
