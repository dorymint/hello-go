// /base/slice/shift/append.
package main

import "fmt"

func main() {
	ss := []string{"hello", "world"}
	printss := func() {
		fmt.Printf("len:%d cap:%d %+q pointer:%p\n", len(ss), cap(ss), ss, ss)
	}

	printss()

	ls := []string{"foo", "bar", "fizz", "buzz"}
	fmt.Printf("shift:%+q\n", ls)
	for _, s := range ls {
		ss = append(ss[1:], s)
		printss()
	}
}
