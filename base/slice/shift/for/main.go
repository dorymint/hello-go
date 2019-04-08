// base/slice/shift/for.
package main

import "fmt"

func Push(ss []string, s string) {
	index := len(ss) - 1
	for i := 0; i < index; i++ {
		ss[i] = ss[i+1]
	}
	ss[index] = s
}

func Printss(ss []string) {
	fmt.Printf("len:%d cap:%d %+q pointer:%p\n", len(ss), cap(ss), ss, ss)
}

func main() {
	ss := []string{"hello", "world", "foo", "bar"}
	Printss(ss)
	Push(ss, "hello world")
	Printss(ss)
}
