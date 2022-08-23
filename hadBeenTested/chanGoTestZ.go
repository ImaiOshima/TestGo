package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ready := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			println("start.")
			<-ready
			println("end.")
		}()
	}
	time.Sleep(time.Second)
	println("======")
	close(ready)
	wg.Wait()
}
