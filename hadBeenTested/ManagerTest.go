package main

import "fmt"

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.age = 12
	m.name = "wyh"
	m.title = "zzr"

	fmt.Println(m)
}
