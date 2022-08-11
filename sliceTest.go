package main

import "fmt"

func main() {
	x:= make([]int,0,5)
	for i:=0;i<8;i++{
		x = append(x,i)
	}
	fmt.Println(x)

	y := make(map[string]int)
	y["a"] = 1
	z,ok := y["a"]
	fmt.Println(z,ok)
	delete(y,"a")
}
