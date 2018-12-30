package main

import (
	"fmt"
)

var ch = make(chan string)

func sendCh() chan<- string {
	return ch
}

func receive() <-chan string {
	return ch
}

func main() {
	sch := sendCh()
	rch := receive()

	go func() {
		sch <- "hello"
	}()
	fmt.Println(<-rch)

	// compile error, allowed only send
	//fmt.Println(<-sch)

	// compile error, allowed only receive
	//rch <- "hi"
}
