package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 1; i++ {
		switch {
		case true:
			fmt.Println("hello")
			break // target is switch
		}
		fmt.Println("world") // reachable
	}

	fmt.Println("----- use label -----")
FOR:
	for i := 0; i < 1; i++ {
		switch {
		case true:
			fmt.Println("hello")
			break FOR // target is FOR:
		}
		fmt.Println("world") // unreachable
	}
}
