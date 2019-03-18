// cap and pointer.
package main

import "fmt"

func main() {
	ss := make([]string, 0, 1)
	add := func(str string) {
		fmt.Printf("append:%q\n", str)
		ss = append(ss, str)
		fmt.Printf("pointer:%p, len:%d, cap:%d\n\n", ss, len(ss), cap(ss))
	}

	fmt.Printf("pointer:%p, len:%d, cap:%d\n\n", ss, len(ss), cap(ss))
	add("hello") // cap:1 pointer not changed
	add("world") // cap:2 pointer changed
	add("foo")   // cap:4 pointer changed
	add("bar")   // cap:4 pointer not changed
}
