// range ch.
//
//   for i := range ch {}
//
package main

import (
	"fmt"
)

func main() {
	ich := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ich <- i
		}
		close(ich)
	}()
	// keep receive until closed
	for i := range ich {
		fmt.Println(i)
	}

	sch := make(chan string, 10)
	sch <- "hello"
	sch <- "world"
	close(sch)
	for v := range sch {
		fmt.Println(v)
	}

	rch := make(chan rune, 10)
	rch <- 'h'
	rch <- 'i'
	for len(rch) != 0 {
		fmt.Println(<-rch)
	}
}
