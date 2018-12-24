package main

import (
	"fmt"
)

func main() {
	var n []string
	fmt.Println("len ", len(n))
	fmt.Println("len ", len([]string{}))
	fmt.Println("len ", len(*new([]string)))
}
