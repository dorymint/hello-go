// base/slice/make.
package main

import "fmt"

type Struct struct{}

func main() {
	ss := make([]*Struct, 10)
	for i := range ss {
		fmt.Println(i)
		if ss[i] != nil {
			panic("want nil")
		}
	}
	fmt.Printf("len:%d cap:%d %p\n", len(ss), cap(ss), ss)
}
