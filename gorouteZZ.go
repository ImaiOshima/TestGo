package main

import "fmt"

type Age uint

func (age Age) String() {
	fmt.Println("this age is ", age)
}

func (age *Age) Modify(value int) {
	*age = Age(value)
}

func main() {
	age := Age(25)
	//page := &age
	mm := (*Age).Modify
	mm(&age, 50)
	age.String()
}
