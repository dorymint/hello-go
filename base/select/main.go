package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan string)
	chd := make(chan int)
	quite := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("print deamon exited")
			wg.Done()
		}()
		for {
			select {
			case s := <-ch:
				fmt.Println(s)
				wg.Done()
			case d := <-chd:
				fmt.Println(d)
				wg.Done()
			case <-quite:
				return
			}
		}
	}()

	wg.Add(1)
	ch <- "hello world"

	wg.Add(1)
	chd <- 1

	wg.Add(1)
	quite <- struct{}{}
	wg.Wait()
}
