package main

import (
	"errors"
	"fmt"
)

func main() {
	test(10,0)
}

func test(a,b int) {
	m := make(map[string]int)

	m["a"] = 1
	x,ok := m["b"]
	fmt.Println(x,ok)
}

func div(a,b int) (int ,error){
	if a == 0 {
		return 0,errors.New("division by zero")
	}
	return a/b,nil
}
