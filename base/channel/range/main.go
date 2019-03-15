// range ch.
//
//   for i := range ch {}
//
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	// keep receive until closed
	for i := range ch {
		fmt.Println(i)
	}
}
