// label for continue.
package main

import "fmt"

func main() {
loop1:
	for i := 0; i < 4; i++ {
		fmt.Println()

	loop2:
		for i := 0; i < 4; i++ {
			switch i {
			case 2:
				continue loop2
			}
			fmt.Printf("loop2: %d\n", i)
		}

		switch i {
		case 2:
			continue loop1
		}
		fmt.Printf("loop1: %d\n", i)
	}
}
