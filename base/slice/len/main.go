// len.
package main

import (
	"fmt"
)

func main() {
	ss := []string{"hello world"}
	printss := func() {
		fmt.Printf("%q\n\tlen:%d\n\n", ss, len(ss))
	}

	printss()

	ss = append(ss, "foo", "bar")
	printss()

	s := "hello world"
	fmt.Printf("max:%s\n", string(s[len(s):]))
}
