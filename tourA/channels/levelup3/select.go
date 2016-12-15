package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Echo is return func point
func Echo() (<-chan string, *sync.WaitGroup, func(string)) {
	ch := make(chan string)
	wg := new(sync.WaitGroup)

	echo := func(s string) {
		defer wg.Done()
		ch <- s
	}
	return ch, wg, echo
}

func msgCat(ch <-chan string, closed <-chan bool) {
	for {
		select {
		case x := <-ch:
			fmt.Println("cat", x)
		case <-closed:
			return
		}
	}
}

func main() {
	fmt.Println("GOROUTINES = ", runtime.NumGoroutine())
	ch, wg, echo := Echo()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go echo(fmt.Sprintf("%8v", i))
	}
	closed := make(chan bool)
	fmt.Println("GOROUTINES = ", runtime.NumGoroutine())
	go func() {
		for {
			select {
			case x := <-ch:
				fmt.Println("dog", x)
			case <-closed:
				return
			}
		}
	}()
	fmt.Println("GOROUTINES = ", runtime.NumGoroutine())
	go msgCat(ch, closed)
	fmt.Println("GOROUTINES = ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("GOROUTINES = ", runtime.NumGoroutine())
	closed<-true
	closed<-true
	fmt.Println("GOROUTINES = ", runtime.NumGoroutine())
}
