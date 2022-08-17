package main

import "fmt"

func main() {
	x := 10
	var p *int = &x
	*p += 20
	fmt.Println(p, *p)
	var p2 *int
	p2 = &x
	fmt.Println(p2 == p)
}
