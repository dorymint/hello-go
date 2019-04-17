// base/slice/copy.
package main

import "fmt"

type Struct struct {
	ss []string
}

func (s *Struct) PopAll() []string {
	var ss []string
	copy(ss, s.ss)
	return ss
}

func main() {
	ss := &Struct{
		ss: []string{"hello", "world"},
	}
	printSS := func(prefix string, ss []string) {
		fmt.Printf("%s%p %+q\n", prefix, ss, ss)
	}

	printSS("ss:", ss.ss)
	newss := ss.PopAll()
	printSS("newss:", newss)
	printSS("ss:", ss.ss)
}
