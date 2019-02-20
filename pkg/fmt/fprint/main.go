package main

import "fmt"

func main() {
	s := "hello world"
	b := []byte("hello world")
	fmt.Printf("s.s: %s\n", s)
	fmt.Printf("s.v: %v\n", s)
	fmt.Printf("b.s: %s\n", b)
	fmt.Printf("b.v: %v\n", b)
}
