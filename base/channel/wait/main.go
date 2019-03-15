// sync.WaitGroup.
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			fmt.Printf("%d loops\n", <-ch)
		}
	}()

	wg.Wait()
}
