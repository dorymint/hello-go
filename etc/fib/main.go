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
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(9))
	fmt.Println(fib(10))

	/// destroy
	fmt.Println(fib(10000000000000000))
}
