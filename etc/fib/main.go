package main

import (
	"fmt"
)

func fib(n uint) uint {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	for i := uint(1); i <= 10; i++ {
		fmt.Printf("i=%d\t%d\n", i, fib(i))
	}

	// stack overflow
	//fmt.Println(fib(10000000000000000))
}
