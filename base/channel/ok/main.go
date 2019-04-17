// base/channel/ok.
package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	for i := 0; i < cap(ch); i++ {
		ch <- i
	}
	close(ch)
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(i, ok)
	}
}
