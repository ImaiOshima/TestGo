package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)
		for {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("data:", x)
			default:
			}
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 5)
	c <- 100
	close(c)
	<-done
}

type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			println("recv:", x)
		}
	}()
	return r
}
func main() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2
	close(r.data)
	r.Wait()
}
