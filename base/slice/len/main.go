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

	fmt.Printf("ss[1:1]:%q len:%d, cap:%d\n\n", ss[1:1], len(ss[1:1]), cap(ss[1:1]))

	s := "hello world"
	fmt.Printf("max:%s\n", string(s[len(s):]))
}
