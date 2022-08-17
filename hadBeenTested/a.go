package main

import "fmt"

func main() {
	a := 1.0 << 3
	fmt.Printf("%T,%v\n", a, a)

	var s uint = 3
	//b:= 1.0 << s
	//fmt.Println("%T,%v",b,b)

	var c int32 = 1.0 << s
	fmt.Printf("%T,%v\n", c, c)
}
