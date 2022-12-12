package main

import "fmt"

type People struct {
}

func (*People) getName(name string) {
	fmt.Println("Hi," + name + ",I am Barry")
}

type Student struct {
	*People
}

func main() {
	name := "Barry"
	a := People{}
	a.getName(name)

	b := Student{&People{}}
	b.getName(name)
}
