// for, switch, break.
package main

import (
	"fmt"
)

func canNotBreakLoop() {
	fmt.Println("--- can not break the for loops ---")
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			fmt.Printf("%d hello\n", i)
		case 4:
			fmt.Printf("%d expected break the loop but can not\n", i)
			break
		default:
			fmt.Println(i)
		}
	}
}

func breakLoop() {
	fmt.Println("--- label break ---")
L:
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			fmt.Printf("%d hello\n", i)
		case 4:
			fmt.Printf("%d break\n", i)
			break L
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	canNotBreakLoop()
	breakLoop()
}
