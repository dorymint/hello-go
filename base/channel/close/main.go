package main

import (
	"fmt"
)

func main() {
	// not locked
	closeCh := make(chan string)
	close(closeCh)
	fmt.Printf("closed channel %v\n", <-closeCh)
}
