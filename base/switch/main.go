package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- can not break the for loops ---")
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			fmt.Println("hello")
		case 4:
			fmt.Printf("can not break!\n")
			fmt.Printf("if you want to stop the for loop then use label\n")
			break
		default:
			fmt.Println(i)
		}
	}

	fmt.Println("--- label break ---")
L:
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			fmt.Println("hello")
		case 4:
			fmt.Println("break")
			break L
		default:
			fmt.Println(i)
		}
	}
}
