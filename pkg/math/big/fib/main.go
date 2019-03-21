// fib.
package main

import (
	"fmt"
	"math/big"
)

func fib(count int) {
	fib := make([]*big.Int, count)
	for i := range fib {
		fib[i] = new(big.Int)
	}
	fib[0].SetInt64(0)
	fib[1].SetInt64(1)
	for i := 2; i < len(fib); i++ {
		fib[i].Add(fib[i-2], fib[i-1])
	}
	for i := 0; i < len(fib); i++ {
		fmt.Println(fib[i])
	}
}

func main() {
	fib(100)
}
