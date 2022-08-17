package main

import "fmt"

func main() {
	const (
		x uint16 = 126
		y
		s = "abc"
		z
	)

	const (
		i = iota
		j
		w
	)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
	fmt.Println(i, j, w)
}
