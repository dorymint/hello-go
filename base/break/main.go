// break on switch statements.
package main

import "fmt"

func main() {
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			fmt.Printf("%d break from switch\n", i)
			break
		case 3:
			fmt.Printf("%d break from loop\n", i)
			break loop
		default:
			fmt.Println(i)
		}
	}
}
