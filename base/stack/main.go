// stack.
package main

import "fmt"

type stack struct {
	vs []interface{}
}

func (s *stack) Push(v interface{}) {
	s.vs = append(s.vs, v)
}

func (s *stack) Pop() interface{} {
	// to do fix for out of bound.
	n := len(s.vs)
	v := s.vs[n-1]
	s.vs = s.vs[:n-1]
	return v
}

func (s *stack) Len() int {
	return len(s.vs)
}

func main() {
	s := new(stack)
	s.Push("hello")
	s.Push("world")
	s.Push(0)
	s.Push(1)

	for s.Len() != 0 {
		v := s.Pop()
		fmt.Printf("%+v\n", v)
	}
}
