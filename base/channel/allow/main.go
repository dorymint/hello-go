// limitation.
package main

import (
	"fmt"
)

var ch = make(chan string)

func sendCh() chan<- string {
	return ch
}

func receiveCh() <-chan string {
	return ch
}

func main() {
	sch := sendCh()
	rch := receiveCh()

	go func() {
		sch <- "hello"
	}()
	fmt.Println(<-rch)

	// compile error, allow only send
	//fmt.Println(<-sch)

	// compile error, allow only receive
	//rch <- "hi"
}
