package main

import (
	"fmt"
)

func main() {
	fmt.Println(0)

	// stack
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
}
