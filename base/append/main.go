package main

import "fmt"

func main() {
	s := []string{"hello", "world"}

	s = append(s, "foo", "bar")
	s = append(s, []string{"fizz", "buzz"}...)

	fmt.Println(s)
}
