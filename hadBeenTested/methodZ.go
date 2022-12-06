package main

import "fmt"

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}

func (n N) value() {
	fmt.Printf("v:%p, %v\n", &n, n)
}

func (n *N) poniter() {
	fmt.Printf("v:%p, %v\n", n, *n)
}

func main() {
	var a N = 25
	//p := &a

	a++
	f1 := a.poniter

	a++
	f2 := a.poniter

	a++
	fmt.Printf("v:%p, %v\n", &a, a)

	f1()
	f2()
}
