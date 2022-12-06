package main

import (
	"fmt"
	"sync"
)

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}

func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))
	for _, ch := range ins {
		go p(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	//bu := buy(10)
	////bu2 := buy(15)
	//pa1 := build(bu)
	//pa2 := build(bu)
	//pa3 := build(bu)
	//pa := merge(pa1,pa2,pa3)
	//thing := pack(pa)
	//for c := range thing {
	//	fmt.Println(c)
	//}
	//var wg sync.WaitGroup
	//test := make(chan string)
	//
	//wg.Add(1)
	//go func(){
	//	defer wg.Done()
	//	for c:= range test {
	//		fmt.Println(c)
	//	}
	//}()
	//time.Sleep(time.Second)
	//test <- "0"
	//test <- "1"
	//test <- "0"
	//close(test)
	////for c := range test {
	////	fmt.Println(c)
	////}
	//
	//wg.Wait()
	sp := new(string)
	fmt.Println(*sp)
	fmt.Println(*sp == "")
}
