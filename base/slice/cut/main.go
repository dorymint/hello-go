// slice.
package main

import "fmt"

func main() {
	i := 0
	lc := func(ss []string) {
		fmt.Printf("%d len:%d cap:%d %+q\n", i, len(ss), cap(ss), ss)
		i++
	}
	ss := []string{"hello", "world"}

	lc(ss)
	lc(ss[0:])
	lc(ss[:0])
}
