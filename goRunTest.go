package main

import (
	"fmt"
	"time"
)

func task(id int){
	for i:=0;i<5;i++ {
		fmt.Printf("%d.--%d\n",id,i)
		time.Sleep(time.Second)
	}
}

func comuser(data chan int,done chan bool){
	for x := range data {
		println("recv:",x)
	}

	done <- true
}

func producer(data chan int){
	for i:=0; i<4; i++ {
		data <- i
	}
	close(data)
}

func main() {
	//go task(1)
	//go task(2)
	//
	//time.Sleep(time.Second * 60)

	data := make(chan int)
	done := make(chan bool)

	go comuser(data,done)
	go producer(data)

	<-done

}
