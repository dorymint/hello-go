package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; true; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := range ch {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
