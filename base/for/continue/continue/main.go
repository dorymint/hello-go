// continue.
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		switch i {
		case 3:
			fmt.Print("hello ")
		case 5:
			continue
		case 9:
			continue
		}
		fmt.Println(i)
	}
}
