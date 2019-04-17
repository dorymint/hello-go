// base/append.
package main

import "fmt"

func main() {
	// append string.
	s := []string{"hello", "world"}
	s = append(s, "foo", "bar")
	s = append(s, []string{"fizz", "buzz"}...)
	fmt.Println(s)

	// check a pointer.
	ss := []string{"hello", "world"}
	printP := func() {
		fmt.Printf("%p %p %+q\n", ss, &ss, ss)
	}
	printP()
	ss = append(ss, "foo")
	printP()
	ss = ss[:0]
	printP()
}
