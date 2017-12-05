package main

import (
	"fmt"
)

func fact(n uint64) uint64 {
	if n > 0 {
		tmp := n* fact(n-1)
		return tmp
	}
	return 1
}

func main() {
	fmt.Println(fact(4))
	fmt.Println(fact(100))
}
