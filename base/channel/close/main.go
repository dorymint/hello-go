// close channel.
package main

import (
	"fmt"
)

func main() {
	// not locked
	closeCh := make(chan string)
	close(closeCh)
	<-closeCh
	fmt.Printf("not locked\n")

	// with buffer
	ch := make(chan int, 30)
	go func() {
		defer func() { fmt.Println("with buffer: close(ch)") }()
		for i := 0; i < 30; i++ {
			ch <- i
		}
		close(ch)
	}()
	for i := range ch {
		fmt.Println("i", i)
	}
}
