// basic usage.
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("sending...")
		time.Sleep(1 * time.Second)
		ch <- "hello world"
	}()
	fmt.Println(<-ch)
}
