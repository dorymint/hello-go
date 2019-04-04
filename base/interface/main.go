package main

import "fmt"

type I interface {
	Name() string
}

type S struct {
	name string
}

func (s *S) Name() string {
	return s.name
}

func main() {
	var i I = &S{"foo"}
	fmt.Println(i.Name())
}
