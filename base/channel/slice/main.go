// send after set nil.
package main

import "fmt"

func main() {
	ss := []string{"hello", "world"}
	ch := make(chan []string, 1)
	ch <- ss
	ss = nil
	fmt.Println(<-ch) // [hello world]
	fmt.Println(ss)   // []
}
