package main

import (
	"fmt"
	"sync"
	"time"
)

var chRoot = make(chan string)

func getch() chan<- string {
	sendCh := make(chan string)
	go func() {
		for {
			chRoot <- <-sendCh
		}
	}()
	return sendCh
}

func main() {
	var wg sync.WaitGroup
	go func() {
		for {
			fmt.Print(<-chRoot)
			wg.Done()
		}
	}()

	ch := getch()
	for i := 0; i < 10; i++ {
		wg.Add(1)

		ch <- fmt.Sprintf("i %d\n", i)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
