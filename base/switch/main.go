package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			fmt.Println("hi!")
		case 4:
			fmt.Printf("\tcan not break!\n\tif you want to stop the for loop then use label\n")
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
			fmt.Println("hi!")
		case 4:
			break L
		default:
			fmt.Println(i)
		}
	}
}
