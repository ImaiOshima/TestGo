package main

import (
	"fmt"
	"sync"
	"time"
)

func watchDog(done chan bool, s string) {
	for {
		select {
		case <-done:
			fmt.Println("停止监控 返回")
			return
		default:
			fmt.Println(s)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)
	s := "watch dog is working"
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchDog(done, s)
	}()
	time.Sleep(2 * time.Second)
	done <- true
	fmt.Println(s)
	wg.Wait()
}
