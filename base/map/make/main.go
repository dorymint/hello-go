package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Printf("%q\n", m)
	fmt.Println("len", len(m))
	fmt.Println("one", m["one"])

	{
		i, ok := m["one"]
		fmt.Printf("i=%d ok=%v\n", i, ok)
	}


	{
		delete(m, "one")
		i, ok := m["one"]
		fmt.Println("after delete key of \"one\"")
		fmt.Printf("i=%d ok=%v\n", i, ok)
		fmt.Println("len", len(m))
	}
}
