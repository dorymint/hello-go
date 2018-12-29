package main

import (
	"context"
	"fmt"
	"time"
)

func graceful() {
	ch := make(chan string)
	topLevelCTX := context.Background()
	ctx, cancel := context.WithCancel(topLevelCTX)
	defer cancel()
	go func() {
		defer fmt.Println("returned goroutine")
		for {
			select {
			case <-ctx.Done():
				return
			case s := <-ch:
				fmt.Println(s)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- fmt.Sprint("hello", i)
		if i == 5 {
			break
		}
	}
}

func main() {
	graceful()
	time.Sleep(time.Second)
}
