// nested continue.
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		switch i {
		case 2:
			switch i {
			case 2:
				continue
			}
		}
		fmt.Println(i)
	}
}
