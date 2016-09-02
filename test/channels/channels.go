package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func lazySum(s []int, c chan int) {
	time.Sleep(5000 * time.Millisecond)
	sum(s, c)
}

func main() {
	s := []int{ 7, 2, 8, -9, 4, 0 }

	// int型のチャンネル
	c := make(chan int)
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)
	go lazySum(s[:], c)

	// <- channel operator
	x, y := <-c, <-c
	lazy := <-c
	// 上のチャンネルが同期するまで次に行かない
	fmt.Println("test1")

	fmt.Println(x, y)
	fmt.Println(lazy)

	fmt.Println("test2")

}

// NOTE:channel operator は goroutine を同期する
