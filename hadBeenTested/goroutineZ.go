package main

import (
	"sync"
	"time"
)

var c int

func counter() int {
	c++
	return c
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Wait()
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done.")
		wg.Done()
	}()
	wg.Wait()
	println("main exit.")
}
