// base/slice/append.
package main

import "fmt"

type Struct struct {
	s string
}

func PrintStructs(ss []*Struct) {
	fmt.Printf("%#v\n", ss)
	for i := range ss {
		fmt.Printf("\t%#v\n", ss[i])
	}
}

func AppendString(ss *[]*Struct, s string) {
	*ss = append(*ss, &Struct{s: s})
}

func main() {
	var ss []*Struct
	PrintStructs(ss)

	ss = append(ss, &Struct{s: "hello"})
	PrintStructs(ss)

	AppendString(&ss, "world")
	PrintStructs(ss)
}
