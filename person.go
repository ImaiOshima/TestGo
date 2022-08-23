package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s,the age is %d", p.name, p.age)
}

func (p *person) test() {
	(*p).name = "zzz"
	fmt.Println("test", p.name)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	p := person{"abc", 20}
	pp := &p
	fmt.Println(pp.name)
	fmt.Println(pp.age)
	printString(p)
	p.test()
}
