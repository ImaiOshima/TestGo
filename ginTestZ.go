package main

import "fmt"

func main() {
	done := make(chan string)
	go func() {
		done <- "a"
		defer fmt.Println("done 已关闭")
		defer close(done)
	}()
	<-done
	fmt.Println("main end")
}
