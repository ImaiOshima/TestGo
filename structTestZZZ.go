package main

import "fmt"

func main() {
	u := struct {
		name string
		age  byte
	}{
		name: "Tom",
		age:  12,
	}
	type file struct {
		name string
		attr struct {
			owner int
			perm  int
		}
	}

	f := file{
		name: "test.dat",
	}
	f.attr.owner = 1
	f.attr.perm = 0755
	fmt.Println(u, f)

	exit := make(chan struct{})

	go func() {
		fmt.Println("hello world")
		exit <- struct{}{}
	}()

	<-exit
	fmt.Println("end.")
}
