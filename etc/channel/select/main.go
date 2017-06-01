package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	ch <- "world"

	mock := make(chan string, 1)
	mock <- "hello"
	var chp *chan string
	chp = &mock

	fmt.Printf("len: *chp=%v, ch=%v\n", len(*chp), len(ch))
	for len(*chp)+len(ch) != 0 {
		select {
		case s := <-*chp:
			fmt.Println("pointa:", s)
		default:
			fmt.Println(<-ch)
		}
	}
}
