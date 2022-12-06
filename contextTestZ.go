package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func watchDog(ctx context.Context, s string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("接受到停止命令 立即返回")
			return
		default:
			fmt.Println(s)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	s := "watch dog is working"
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchDog(ctx, s)
	}()
	time.Sleep(5 * time.Second)
	stop()
	fmt.Println("stop结束")
	wg.Wait()
}
