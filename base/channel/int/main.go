// chan int.
package main

import (
	"fmt"
	"time"
)

func main() {
	ci := make(chan int)      // unbuffered
	cj := make(chan int, 0)   // unbuffered
	ck := make(chan int, 100) // buffered

	go func() {
		fmt.Println("send to ci")
		ci <- 1
	}()
	fmt.Println(<-ci)
	fmt.Printf("ci done\n\n")

	go func() {
		fmt.Println("send to cj")
		time.Sleep(3 * time.Second)
		cj <- 2
	}()
	fmt.Println("blocked, wait 3 seconds")
	fmt.Println(<-cj)
	fmt.Printf("cj done\n\n")

	fmt.Println("send to ck")
	ck <- 3 // non blocking
	fmt.Println(<-ck)
	fmt.Printf("ck done\n\n")
}
