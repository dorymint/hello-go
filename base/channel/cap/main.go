// capacity.
package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// not blocked
	ch <- 0

	fmt.Printf("cap %+v", cap(ch))
}
