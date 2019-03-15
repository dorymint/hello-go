// delete in map.
//
//	delete(Map, Key)
//
package main

import "fmt"

func main() {
	m := map[string]int{
		"hello":  0,
		"gopher": 1,
		"world":  2,
	}

	fmt.Printf("before:%+v\n", m)

	delete(m, "gopher")

	fmt.Printf("after:%+v\n", m)
}
