package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() {
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			println(name, x)
		}

	}()

	go func() {
		defer wg.Done()
		defer close(b)
		defer close(a)
		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i * 10:
			}
		}
	}()

	//c := make(chan int)
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	x, ok := <-c
	//	println(x, ok)
	//}()
	wg.Wait()

}
