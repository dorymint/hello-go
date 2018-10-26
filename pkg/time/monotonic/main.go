package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	for {
		now := time.Now()
		fmt.Println(now, now.Sub(t))
		time.Sleep(time.Second)
	}
}
