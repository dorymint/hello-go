// reset.
package main

import "fmt"

func main() {
	ss := []string{"hello", "world"}
	ss = ss[:0]
	fmt.Printf("len:%d cap:%d %+q\n", len(ss), cap(ss), ss)
}
