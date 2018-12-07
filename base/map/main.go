package main

import (
	"fmt"
)

func forMap(m map[int]bool) {
	c := 0
	for key, b := range m {
		fmt.Printf("%#v=%q\n", key, b)
		c++
	}
	fmt.Printf("%d loops\n", c)
}

func main() {
	m := make(map[string]bool)
	const key = "key strings"
	m[key] = true

	fmt.Printf("%#v=%v len=%d\n", key, m[key], len(m))
	delete(m, key)
	fmt.Printf("%#v=%v len=%d\n", key, m[key], len(m))

	// map is unordered access
	mint := make(map[int]bool)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			mint[i] = true
		} else {
			mint[i] = false
		}
	}
	forMap(mint)
}
