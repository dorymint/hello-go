// base/make/cap.
package main

import "fmt"

type Struct struct{ s string }

func main() {
	ss := make([]*Struct, 0, 10)
	fmt.Println("ss == nil:", ss == nil) // false

	sss := []*Struct{{"hello"}}
	sss = append(sss, ss...)
	fmt.Printf("sss: len:%d cap:%d %+v\n", len(sss), cap(sss), sss)
}
