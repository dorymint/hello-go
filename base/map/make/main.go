// make map.
//
//	newMap := make(map[Tyep]Type)
//	newMapWithSpace := make(map[Tyep]Type, NumberOfElements)
//
package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	printLen := func() func() {
		var n int
		return func() {
			fmt.Printf("%d:\n", n)
			fmt.Printf("  %+v:\n", m)
			fmt.Printf("  Len:%d\n\n", len(m))
			n++
		}
	}()

	m = make(map[string]int)
	printLen()

	// with number of elements
	m = make(map[string]int, 100)
	printLen() // not different output

	m = map[string]int{}
	printLen()
}
