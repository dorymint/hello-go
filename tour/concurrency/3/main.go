// Buffered Channels.
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// deadlock.
	for i := 0; i <= cap(ch); i++ {
		ch <- i
	}
}
